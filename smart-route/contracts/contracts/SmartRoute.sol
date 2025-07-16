// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title SmartRoute
 * @dev 智能路由聚合器合约 - 可升级版本
 */
contract SmartRoute is 
    Initializable, 
    OwnableUpgradeable, 
    ReentrancyGuardUpgradeable, 
    PausableUpgradeable,
    UUPSUpgradeable 
{
    using SafeERC20 for IERC20;

    // 事件
    event RouteExecuted(
        address indexed user,
        address indexed tokenIn,
        address indexed tokenOut,
        uint256 amountIn,
        uint256 amountOut,
        bytes32 routeId
    );

    event RouteRegistered(
        bytes32 indexed routeId,
        address indexed aggregator,
        string name,
        bool isActive
    );

    event ProtocolFeeUpdated(uint256 oldFee, uint256 newFee);
    event RouteFeeUpdated(bytes32 indexed routeId, uint256 oldFee, uint256 newFee);

    // 路由信息结构
    struct Route {
        address aggregator;
        string name;
        bool isActive;
        uint256 fee; // 手续费，以基点为单位 (1 = 0.01%)
    }

    // 状态变量
    mapping(bytes32 => Route) public routes;
    mapping(address => bool) public authorizedAggregators;
    
    uint256 public totalRoutes;
    uint256 public protocolFee; // 协议手续费

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @dev 初始化函数
     * @param _owner 合约所有者
     * @param _protocolFee 初始协议手续费
     */
    function initialize(address _owner, uint256 _protocolFee) public initializer {
        __Ownable_init(_owner);
        __ReentrancyGuard_init();
        __Pausable_init();
        __UUPSUpgradeable_init();
        
        protocolFee = _protocolFee;
    }

    /**
     * @dev 注册新的路由
     * @param routeId 路由唯一标识
     * @param aggregator 聚合器地址
     * @param name 路由名称
     * @param fee 手续费
     */
    function registerRoute(
        bytes32 routeId,
        address aggregator,
        string memory name,
        uint256 fee
    ) external onlyOwner {
        require(aggregator != address(0), "Invalid aggregator address");
        require(routes[routeId].aggregator == address(0), "Route already exists");
        require(fee <= 1000, "Fee too high"); // 最大 10%

        routes[routeId] = Route({
            aggregator: aggregator,
            name: name,
            isActive: true,
            fee: fee
        });

        authorizedAggregators[aggregator] = true;
        totalRoutes++;

        emit RouteRegistered(routeId, aggregator, name, true);
    }

    /**
     * @dev 执行路由交换
     * @param routeId 路由ID
     * @param tokenIn 输入代币地址
     * @param tokenOut 输出代币地址
     * @param amountIn 输入金额
     * @param minAmountOut 最小输出金额
     * @param data 路由特定数据
     */
    function executeRoute(
        bytes32 routeId,
        address tokenIn,
        address tokenOut,
        uint256 amountIn,
        uint256 minAmountOut,
        bytes calldata data
    ) external nonReentrant whenNotPaused returns (uint256 amountOut) {
        Route storage route = routes[routeId];
        require(route.isActive, "Route is not active");
        require(authorizedAggregators[route.aggregator], "Unauthorized aggregator");
        require(amountIn > 0, "Amount must be greater than 0");

        // 转移代币到聚合器
        IERC20(tokenIn).safeTransferFrom(msg.sender, route.aggregator, amountIn);

        // 调用聚合器执行交换
        (bool success, bytes memory result) = route.aggregator.call(data);
        require(success, "Route execution failed");

        // 解析输出金额
        amountOut = abi.decode(result, (uint256));
        require(amountOut >= minAmountOut, "Insufficient output amount");

        // 计算手续费
        uint256 protocolFeeAmount = (amountOut * protocolFee) / 10000;
        uint256 routeFeeAmount = (amountOut * route.fee) / 10000;
        uint256 userAmountOut = amountOut - protocolFeeAmount - routeFeeAmount;

        // 转移代币给用户
        IERC20(tokenOut).safeTransferFrom(route.aggregator, msg.sender, userAmountOut);

        emit RouteExecuted(msg.sender, tokenIn, tokenOut, amountIn, amountOut, routeId);
    }

    /**
     * @dev 更新路由状态
     */
    function updateRouteStatus(bytes32 routeId, bool isActive) external onlyOwner {
        require(routes[routeId].aggregator != address(0), "Route does not exist");
        routes[routeId].isActive = isActive;
    }

    /**
     * @dev 更新路由手续费
     */
    function updateRouteFee(bytes32 routeId, uint256 newFee) external onlyOwner {
        require(routes[routeId].aggregator != address(0), "Route does not exist");
        require(newFee <= 1000, "Fee too high"); // 最大 10%
        
        uint256 oldFee = routes[routeId].fee;
        routes[routeId].fee = newFee;
        
        emit RouteFeeUpdated(routeId, oldFee, newFee);
    }

    /**
     * @dev 更新协议手续费
     */
    function updateProtocolFee(uint256 newFee) external onlyOwner {
        require(newFee <= 100, "Fee too high"); // 最大 1%
        uint256 oldFee = protocolFee;
        protocolFee = newFee;
        emit ProtocolFeeUpdated(oldFee, newFee);
    }

    /**
     * @dev 暂停合约
     */
    function pause() external onlyOwner {
        _pause();
    }

    /**
     * @dev 恢复合约
     */
    function unpause() external onlyOwner {
        _unpause();
    }

    /**
     * @dev 紧急提取代币
     */
    function emergencyWithdraw(address token, address to) external onlyOwner {
        uint256 balance = IERC20(token).balanceOf(address(this));
        IERC20(token).safeTransfer(to, balance);
    }

    /**
     * @dev 升级合约实现
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    /**
     * @dev 获取路由信息
     */
    function getRoute(bytes32 routeId) external view returns (
        address aggregator,
        string memory name,
        bool isActive,
        uint256 fee
    ) {
        Route storage route = routes[routeId];
        return (route.aggregator, route.name, route.isActive, route.fee);
    }

    /**
     * @dev 获取路由总数
     */
    function getTotalRoutes() external view returns (uint256) {
        return totalRoutes;
    }
} 