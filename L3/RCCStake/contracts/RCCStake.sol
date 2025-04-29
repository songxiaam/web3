// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/utils/math/Math.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

// import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";

// 质押 = 存款
// 解除质押 = 赎回

contract RCCStake is
    Initializable,
    UUPSUpgradeable,
    PausableUpgradeable,
    AccessControlUpgradeable
{
    using SafeERC20 for IERC20;
    using Address for address;
    using Math for uint256;
    using Strings for uint256;

    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyRole(UPGRADE_ROLE) {}

    // 质押池
    // 一个token = 10^18
    struct Pool {
        address stTokenAddress; //质押代币地址, 一种币类型
        uint256 poolWeight; //质押池权重
        uint256 lastRewardBlock; //上次更新奖励区块
        uint256 accRCCPerST; //单个质押代币累计的RCC奖励, 保证精度, 乘了一个进度因子 10^18,accRCCPerST*1e18=奖励的token
        uint256 stTokenAmount; //总质押代币量
        uint256 minDepositAmount; //最小质押金额
        uint256 unstakeLockedBlocks; //解除质押的锁定区块数, 用户发起接触质押后,需要等待多少个区块才能真正提取质押资产
    }
    // 用户质押信息
    struct User {
        uint256 stAmount; //用户质押代币量
        uint256 finishedRCC; //已分配的RCC数量, 可以理解为已结算的RCC, 从这个值开始计算新的奖励,每次质押或解质押时,重新计算当前代币奖励价值,
        uint256 pendingRCC; //待领取的RCC数量, claim时取出,计算pendingRCC-finishedRCC的差值
        // Withdraw request list
        UnstakeRequest[] requests;
    }

    // 解质押请求
    struct UnstakeRequest {
        uint256 amount;
        uint256 unlockBlocks;
    }

    // 常量
    bytes32 public constant ADMIN_ROLE = keccak256("admin_role");
    bytes32 public constant UPGRADE_ROLE = keccak256("upgrade_role");
    // 原生币质押池ID
    uint256 public constant nativeCurrency_PID = 0;

    // 质押池, 第一个pool是原生币池
    Pool[] public pools;
    // 用户 (id => (address => User))
    mapping(uint256 => mapping(address => User)) public users;
    // First block that RCCStake will start from
    uint256 public startBlock;
    // First block that RCCStake will end from
    uint256 public endBlock;
    // Total pool weight / Sum of all pool weights
    uint256 public totalPoolWeight;
    // RCC token reward per block 每个区块的RCC奖励
    uint256 public RCCPerBlock;

    // 是否可提现
    bool public canWithdraw;
    // 是否可领取奖励
    bool public canClaim;

    // RCC token
    IERC20 public RCC;

    // 事件
    event AddPool(
        uint256 pid,
        address indexed stTokenAddress,
        uint256 indexed poolWeight,
        uint256 indexed lastRewardBlock,
        uint256 minDepositAmount,
        uint256 unstakeLockedBlocks
    );
    event UpdatePool(
        uint256 indexed poolId,
        uint256 indexed lastRewardBlock,
        uint256 totalRCC
    );

    event Deposit(address sender, uint256 pid, uint256 amount);
    event SetRCC(IERC20 RCC);
    event SetPoolWeight(uint256 indexed pid, uint256 poolWeight);
    event UpdateMinDepositAmount(uint256 pid, uint256 minDepositAmount);
    event UpdateUnstakeLockBlocks(uint256 pid, uint256 unstakeLockBlocks);
    event PauseWithdraw();
    event ResumeWithdraw();
    event PauseClaim();
    event ResumeClaim();
    event SetStartBlock(uint256 startBlock);
    event SetEndBlock(uint256 endBlock);
    event SetRCCPerBlock(uint256 RCCPerBlock);
    event RequestUnstake(
        address indexed user,
        uint256 indexed pid,
        uint256 amount
    );
    event Withdraw(
        address indexed user,
        uint256 indexed pid,
        uint256 amount,
        uint256 blockNumber
    );
    event Claim(address indexed user, uint256 indexed pid, uint256 amount);


    modifier checkPid(uint256 _pid) {
      require(_pid < pools.length, string(abi.encodePacked("Invalid pool id: ", _pid.toString())));
      _;
    }

    modifier whenCanClaim() {
      require(canClaim, "Claim is paused");
      _;
    }
    modifier whenCanWithdraw() {
      require(canWithdraw, "Withdraw is paused");
      _;
    }

    // 防止重复初始化，只允许执行一次。
    function initialize(
        IERC20 _RCC,
        uint256 _startBlock,
        uint256 _endBlock,
        uint256 _RCCPerBlock
    ) public initializer {
        require(
            _startBlock < _endBlock && _RCCPerBlock > 0,
            "Invalid startBlock or endBlock"
        );

        __AccessControl_init(); //初始化访问控制
        __UUPSUpgradeable_init(); //初始化升级
        __Pausable_init();
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(ADMIN_ROLE, msg.sender);
        _grantRole(UPGRADE_ROLE, msg.sender);

        setRCC(_RCC);
        startBlock = _startBlock;
        endBlock = _endBlock;
        RCCPerBlock = _RCCPerBlock;
    }

    // ============================ Admin Function ============================

    function setRCC(IERC20 _RCC) public onlyRole(ADMIN_ROLE) {
        RCC = _RCC;
        emit SetRCC(_RCC);
    }

    // 创建/更新质押池
    function addPool(
        address _stTokenAddress,
        uint256 _poolWeight,
        uint256 _minDepositAmount,
        uint256 _unstakeLockBlocks,
        bool _withUpdate
    ) public onlyRole(ADMIN_ROLE) {
        if (pools.length > 0) {
            require(_stTokenAddress != address(0x0), "Invalid stTokenAddress");
        } else {
            require(_stTokenAddress == address(0x0), "Invalid stTokenAddress");
        }
        require(_minDepositAmount > 0, "Invalid minDepositAmount");
        require(_unstakeLockBlocks > 0, "Invalid unstakeLockBlocks");

        require(block.number < endBlock, "Alredy ended");
        if (_withUpdate) {
            // 更新
            massUpdatePools();
        }
        uint256 lastRewardBlock = block.number > startBlock
            ? block.number
            : startBlock;

        totalPoolWeight += _poolWeight;
        pools.push(
            Pool({
                stTokenAddress: _stTokenAddress,
                poolWeight: _poolWeight,
                lastRewardBlock: lastRewardBlock,
                accRCCPerST: 0,
                stTokenAmount: 0,
                minDepositAmount: _minDepositAmount,
                unstakeLockedBlocks: _unstakeLockBlocks
            })
        );

        // 更新质押池权重
        emit AddPool(
            pools.length - 1,
            _stTokenAddress,
            _poolWeight,
            lastRewardBlock,
            _minDepositAmount,
            _unstakeLockBlocks
        );
    }

    // 更改质押池属性
    function updateMinDepositAmount(
        uint256 _pid,
        uint256 _minDepositAmount
    ) public onlyRole(ADMIN_ROLE) checkPid(_pid) {
        require(_pid < pools.length, "Invalid pool id");
        Pool storage pool = pools[_pid];
        require(_minDepositAmount > 0, "Invalid minDepositAmount");
        pool.minDepositAmount = _minDepositAmount;
        emit UpdateMinDepositAmount(_pid, _minDepositAmount);
    }

    function updateUnstakeLockBlocks(
        uint256 _pid,
        uint256 _unstakeLockBlocks
    ) public onlyRole(ADMIN_ROLE) checkPid(_pid) {
        require(_pid < pools.length, "Invalid pool id");
        Pool storage pool = pools[_pid];
        require(_unstakeLockBlocks > 0, "Invalid unstakeLockBlocks");
        pool.unstakeLockedBlocks = _unstakeLockBlocks;
        emit UpdateUnstakeLockBlocks(_pid, _unstakeLockBlocks);
    }

    // 修改权重
    function setPoolWeight(
        uint256 _pid,
        uint256 _poolWeight,
        bool _withUpdate
    ) public onlyRole(ADMIN_ROLE) checkPid(_pid) {
        require(_poolWeight > 0, "Invalid pool weight");
        // 再修改前更新所有质押池,更新权重会影响奖励分配,先将未分配奖励分配完再修改
        if (_withUpdate) {
            massUpdatePools();
        }
        pools[_pid].poolWeight = _poolWeight;
        emit SetPoolWeight(_pid, _poolWeight);
    }

    // 暫停提現
    function pauseWithdraw() public onlyRole(ADMIN_ROLE) {
        canWithdraw = false;
        emit PauseWithdraw();
    }

    // 恢復提現
    function resumeWithdraw() public onlyRole(ADMIN_ROLE) {
        canWithdraw = true;
        emit ResumeWithdraw();
    }

    // 暫停領取獎勵
    function pauseClaim() public onlyRole(ADMIN_ROLE) {
        canClaim = false;
        emit PauseClaim();
    }

    // 恢复领取奖励
    function resumeClaim() public onlyRole(ADMIN_ROLE) {
        canClaim = true;
        emit ResumeClaim();
    }

    function setStartBlock(uint256 _startBlock) public onlyRole(ADMIN_ROLE) {
        startBlock = _startBlock;
        emit SetStartBlock(_startBlock);
    }

    function setEndBlock(uint256 _endBlock) public onlyRole(ADMIN_ROLE) {
        endBlock = _endBlock;
        emit SetEndBlock(_endBlock);
    }

    function setRcPerBlock(uint256 _RCCPerBlock) public onlyRole(ADMIN_ROLE) {
        RCCPerBlock = _RCCPerBlock;
        emit SetRCCPerBlock(_RCCPerBlock);
    }

    // ========================================================
    // 更新所有pool
    function massUpdatePools() private {
        uint256 length = pools.length;
        for (uint256 pid = 0; pid < length; pid++) {
            updatePool(pid);
        }
    }

    // 更新质押池,用户更新奖励
    function updatePool(uint256 _pid) public checkPid(_pid) {
        Pool storage pool = pools[_pid];
        // 如果当前区块小于质押池的区块,则直接返回
        if (block.number <= pool.lastRewardBlock) {
            return;
        }

        // 如果没有质押代币,则直接返回
        uint256 stSupply = pool.stTokenAmount;
        if (stSupply == 0) {
            pool.lastRewardBlock = block.number;
            return;
        }

        uint256 rewardAmount = getRewardAmount(
            pool.lastRewardBlock,
            block.number
        );
        // 计算这个pool的可分配总奖励
        uint256 RCCReward = (rewardAmount * pool.poolWeight) / totalPoolWeight;
        // 更新单个代币累计RCC数量
        pool.accRCCPerST += (RCCReward * 1e18) / stSupply;
        // 更新最后奖励区块
        pool.lastRewardBlock = block.number;

        emit UpdatePool(_pid, pool.lastRewardBlock, RCCReward);
    }

    // =========== Query Function ============

    // 计算区块间总奖励, from 到 to 区块的RCC奖励
    function getRewardAmount(
        uint256 _from,
        uint256 _to
    ) public view returns (uint256 multiplier) {
        require(_from <= _to, "Invalid block range");
        if (_from < startBlock) {
            _from = startBlock;
        }
        if (_to > endBlock) {
            _to = endBlock;
        }
        require(_from <= _to, "Invalid block range");
        bool success;
        (success, multiplier) = Math.tryMul(_to - _from, RCCPerBlock);
        require(success, "Multiplier overflow");
    }

    function poolLength() external view returns (uint256) {
        return pools.length;
    }

    // 查询用户在某个pool中待领取奖励
    function pendingRCC(
        uint256 _pid,
        address _user
    ) external checkPid(_pid) view returns (uint256) {
        return pendingRCCByBlockNumber(_pid, _user, block.number);
    }

    function pendingRCCByBlockNumber(
        uint256 _pid,
        address _user,
        uint256 _blockNumber
    ) public checkPid(_pid) view returns (uint256) {
        Pool storage pool = pools[_pid];
        User storage user = users[_pid][_user];

        uint256 accRCCPerST = pool.accRCCPerST;
        uint256 stSupply = pool.stTokenAmount;
        // 最新的accRCCPerST
        if (_blockNumber > pool.lastRewardBlock && stSupply != 0) {
            uint256 rewardAmount = getRewardAmount(
                pool.lastRewardBlock,
                _blockNumber
            );
            // 当前区块分到的奖励
            uint256 RCCForPool = (rewardAmount * pool.poolWeight) /
                totalPoolWeight;
            accRCCPerST = accRCCPerST + (RCCForPool * 1 ether) / stSupply;
        }
        return
            ((user.stAmount * accRCCPerST) / 1 ether - user.finishedRCC) +
            user.pendingRCC;
    }

    // 获取staking amount
    function stakingBalance(uint256 _pid, address _user) external checkPid(_pid) view returns (uint256) {
      return users[_pid][_user].stAmount;
    }

    function withdrawAmount(uint256 _pid, address _user) external checkPid(_pid) view returns (uint256 requestAmount, uint256 pendingAmount) {
      User storage user = users[_pid][_user];
      for (uint i = 0; i < user.requests.length; i++) {
        if (user.requests[i].unlockBlocks <= block.number) {
          pendingAmount += user.requests[i].amount;
        }
        requestAmount += user.requests[i].amount;
      }
    }

    function userUnstakeRequest(uint256 _pid, address _user, uint256 index) external checkPid(_pid) view returns (UnstakeRequest memory) {
      return users[_pid][_user].requests[index];
    }

    function userUnstakeRequestLength(uint256 _pid, address _user) external checkPid(_pid) view returns (uint256) {
      return users[_pid][_user].requests.length;
    }

    function getTotalPoolWeight() external view returns (uint256) {
      return totalPoolWeight;
    }

    // =================== User Function 用户功能 ===================
    // 质押原生币 如ETH, BNB...
    // 原生币直接通过msg.value完成
    function depositNativeCurrency() public whenCanWithdraw() payable {
        Pool storage pool = pools[nativeCurrency_PID];
        require(pool.stTokenAddress == address(0), "Invalid pool");
        uint256 _amount = msg.value;
        require(_amount > pool.minDepositAmount, "deposit amount is too small");
        _deposit(nativeCurrency_PID, _amount);
    }

    // 质押其他ERC20代币
    function deposit(uint256 _pid, uint256 _amount) public checkPid(_pid) {
        Pool storage pool = pools[_pid];
        require(_amount > pool.minDepositAmount, "deposit amount is too small");
        if (_amount > 0) {
            // 先从msg.sender手里把_amount数量的代币安全的转到此合约address(this)中
            IERC20(pool.stTokenAddress).safeTransferFrom(
                msg.sender,
                address(this),
                _amount
            );
        }
        // 处理此合约中状态
        _deposit(_pid, _amount);
    }

    // 用户申请解除质押
    function unstake(uint256 _pid, uint256 _amount) public whenCanWithdraw() checkPid(_pid) {
        Pool storage pool = pools[_pid];
        User storage user = users[_pid][msg.sender];
        require(user.stAmount >= _amount, "Invalid amount");
        updatePool(_pid);

        // 计算用户应得奖励
        uint256 pendingRCC_ = (user.stAmount * pool.accRCCPerST) /
            1 ether -
            user.finishedRCC;
        if (pendingRCC_ > 0) {
            user.pendingRCC = user.pendingRCC + pendingRCC_;
        }
        if (_amount > 0) {
            user.stAmount = user.stAmount - _amount;
            user.requests.push(
                UnstakeRequest({
                    amount: _amount,
                    unlockBlocks: block.number + pool.unstakeLockedBlocks
                })
            );
        }
        pool.stTokenAmount = pool.stTokenAmount - _amount;
        //计算finishedRCC
        user.finishedRCC = (user.stAmount * pool.accRCCPerST) / 1 ether;
        emit RequestUnstake(msg.sender, _pid, _amount);
    }

    // 提现
    function withdraw(uint256 _pid) public whenCanWithdraw() checkPid(_pid) {
        Pool storage pool = pools[_pid];
        User storage user = users[_pid][msg.sender];
        uint256 pendingWithdraw; //可解锁金额
        uint256 popNum; // 可解锁笔数
        for (uint256 i = 0; i < user.requests.length; i++) {
            if (user.requests[i].unlockBlocks > block.number) {
                break;
            }
            pendingWithdraw += user.requests[i].amount;
            popNum++;
        }
        for (uint256 i = 0; i < user.requests.length - popNum; i++) {
            user.requests[i] = user.requests[i + popNum];
        }
        for (uint i = 0; i < popNum; i++) {
            user.requests.pop();
        }
        // Storage arrays do not support slicing operations.
        // user.requests = user.requests[popNum:]

        if (pendingWithdraw > 0) {
            if (pool.stTokenAddress == address(0x0)) {
                // 原生币提现
                _safeNativeCurrencyTransfer(msg.sender, pendingWithdraw);
            } else {
                // 其他代币提现
                IERC20(pool.stTokenAddress).safeTransfer(
                    msg.sender,
                    pendingWithdraw
                );
            }
        }
        emit Withdraw(msg.sender, _pid, pendingWithdraw, block.number);
    }

    // 领取奖励
    function claim(uint256 _pid) public whenCanClaim() checkPid(_pid) {
        Pool storage pool = pools[_pid];
        User storage user = users[_pid][msg.sender];
        // 更新质押池
        updatePool(_pid);
        // 计算应得奖励 距离上一次计算奖励以来,新增的奖励+上次计算应得的奖励
        uint256 pendingRCC_ = (user.stAmount * pool.accRCCPerST) /
            1 ether -
            user.finishedRCC +
            user.pendingRCC;
        // 更新用户质押信息
        if (pendingRCC_ > 0) {
            user.pendingRCC = 0;
            _safeRCCTransfer(msg.sender, pendingRCC_);
        }
        user.finishedRCC = (user.stAmount * pool.accRCCPerST) / 1 ether;
        emit Claim(msg.sender, _pid, pendingRCC_);
    }

    // =============================================================

    // ======== Internal Function ========
    function _deposit(uint256 _pid, uint256 _amount) internal {
        // 获取质押池
        Pool storage pool = pools[_pid];
        // 获取该质押池对应的用户, 不同质押池中可能存在相同address用户
        User storage user = users[_pid][msg.sender];
        // 在质押前把之前的奖励先结算好,否则最近质押的代币也会分配之前的奖励,对于原先的代币不公平
        updatePool(_pid);

        //原先是否已有质押代币, 有的话需要计算应得但未得的奖励
        if (user.stAmount > 0) {
            //计算应得总奖励
            (bool success, uint256 accST) = Math.tryMul(
                user.stAmount,
                pool.accRCCPerST
            );
            require(success, "Multiplier overflow");
            // 单位换算
            (success, accST) = Math.tryDiv(accST, 1 ether);
            require(success, "Multiplier overflow");

            // 本次新增的待领取奖励
            (bool success2, uint256 pendingRCC_) = Math.trySub(
                accST,
                user.finishedRCC
            );
            require(success2, "Multiplier overflow");

            // 新增的pendingRCC 累加到user.pendingRCC
            if (pendingRCC_ > 0) {
                (bool success3, uint256 _pendingRCC) = Math.tryAdd(
                    user.pendingRCC,
                    pendingRCC_
                );
                require(success3, "user pendingRCC overflow");
                user.pendingRCC = _pendingRCC;
            }
        }
        // 更新用户质押代币数量
        if (_amount > 0) {
            (bool success4, uint256 _stAmount) = Math.tryAdd(
                user.stAmount,
                _amount
            );
            require(success4, "user stAmount overflow");
            user.stAmount = _stAmount;
        }

        // 更新质押池质押代币数量
        (bool success5, uint256 stTokenAmount) = Math.tryAdd(
            pool.stTokenAmount,
            _amount
        );
        require(success5, "pool stTokenAmount overflow");
        pool.stTokenAmount = stTokenAmount;

        // 更新用户已领取奖励
        (bool success6, uint256 finishedRCC) = Math.tryMul(
            pool.accRCCPerST,
            user.stAmount
        );
        require(success6, "Multiplier overflow");
        // 单位换算
        (success6, finishedRCC) = Math.tryDiv(finishedRCC, 1 ether);
        require(success6, "Multiplier overflow");
        user.finishedRCC = finishedRCC;

        emit Deposit(msg.sender, _pid, _amount);
    }

    // 安全转账方法
    function _safeRCCTransfer(address _to, uint256 _amount) internal {
        uint256 RCCBalance = RCC.balanceOf(address(this));
        if (_amount > RCCBalance) {
            RCC.transfer(_to, RCCBalance);
        } else {
            RCC.transfer(_to, _amount);
        }
    }

    function _safeNativeCurrencyTransfer(
        address _to,
        uint256 _amount
    ) internal {
        (bool success, bytes memory data) = address(_to).call{value: _amount}(
            ""
        );
        require(success, "Transfer failed");
        if (data.length > 0) {
            success = abi.decode(data, (bool));
            require(success, "Transfer failed");
        }
    }
}
