// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./SmartRoute.sol";

/**
 * @title SmartRouteV2
 * @dev 智能路由聚合器合约 V2 - 升级版本示例
 */
contract SmartRouteV2 is SmartRoute {
    // 新增状态变量
    mapping(address => uint256) public userTradeCount;
    mapping(address => uint256) public userTotalVolume;
    
    // 新增事件
    event UserStatsUpdated(
        address indexed user,
        uint256 tradeCount,
        uint256 totalVolume
    );

    /**
     * @dev 升级后的初始化函数
     * @param _owner 合约所有者
     * @param _protocolFee 初始协议手续费
     */
    function initializeV2(address _owner, uint256 _protocolFee) public reinitializer(2) {
        // 可以在这里初始化 V2 的新状态变量
        // 注意：不要重复初始化 V1 的状态变量
    }

    /**
     * @dev 重写执行路由函数，添加用户统计
     */
    function executeRoute(
        bytes32 routeId,
        address tokenIn,
        address tokenOut,
        uint256 amountIn,
        uint256 minAmountOut,
        bytes calldata data
    ) external nonReentrant whenNotPaused returns (uint256 amountOut) {
        // 调用父合约的执行逻辑
        amountOut = super.executeRoute(routeId, tokenIn, tokenOut, amountIn, minAmountOut, data);
        
        // V2 新增：更新用户统计
        userTradeCount[msg.sender]++;
        userTotalVolume[msg.sender] += amountIn;
        
        emit UserStatsUpdated(msg.sender, userTradeCount[msg.sender], userTotalVolume[msg.sender]);
    }

    /**
     * @dev 获取用户统计信息
     */
    function getUserStats(address user) external view returns (uint256 tradeCount, uint256 totalVolume) {
        return (userTradeCount[user], userTotalVolume[user]);
    }

    /**
     * @dev 批量获取路由信息
     */
    function getMultipleRoutes(bytes32[] calldata routeIds) external view returns (
        address[] memory aggregators,
        string[] memory names,
        bool[] memory isActive,
        uint256[] memory fees
    ) {
        uint256 length = routeIds.length;
        aggregators = new address[](length);
        names = new string[](length);
        isActive = new bool[](length);
        fees = new uint256[](length);
        
        for (uint256 i = 0; i < length; i++) {
            Route storage route = routes[routeIds[i]];
            aggregators[i] = route.aggregator;
            names[i] = route.name;
            isActive[i] = route.isActive;
            fees[i] = route.fee;
        }
    }
} 