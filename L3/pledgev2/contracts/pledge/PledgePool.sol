// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "../library/SafeTransfer.sol";
import "../multiSignature/multiSignatureClient.sol";
import "../interface/IBscPledgeOracle.sol";
import "../interface/IDebtToken.sol";
import "../interface/IUniswapV2Router02.sol";

contract PledgePool is ReentrancyGuard, SafeTransfer, multiSignatureClient {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    uint256 internal constant calDecimal = 1e18;
    uint256 internal constant baseDecimal = 1e8;
    uint256 public minAmount = 100e6;
    uint256 constant baseYear = 365 days;

    enum PoolState {
        MATCH, //撮合中
        EXECUTION,
        FINISH,
        LIQUIDATION,
        UNDONE
    }

    PoolState constant defaultChoice = PoolState.MATCH;

    bool public globalPause = false;

    address public swapRouter;
    address payable public feeAddress; //平台费用转入这个地址
    IBscPledgeOracle public oracle;

    uint256 public lendFee; // 平台收取的利率
    uint256 public borrowFee; // 归还借款人质押borrowToken时需要支付的费率,给平台

    // 每个池基本信息(数据结构)
    struct PoolBaseInfo {
        uint256 settleTime; //结算时间
        uint256 endTime; //结束时间
        uint256 interestRate; //固定利率, 单位 1e8
        uint256 maxSupply; //池最大限额
        uint256 lendSupply; //当前实际存款
        uint256 borrowSupply; //当前实际借款(借款人作为质押的token)
        uint256 mortgageRate; //池的抵押率, 单位1e8
        address lendToken; //贷款代币地址,出借人
        address borrowToken; //借款代币地址
        PoolState state; //当前状态
        IDebtToken spCoin; // sp_token 的erc20地址
        IDebtToken jpCoin; // sp_token 的erc20地址
        uint256 autoLiquidateThreshold; //自动清算阈值
    }

    PoolBaseInfo[] public poolBaseInfo;

    // 每个池的数据信息
    struct PoolDataInfo {
        uint256 settleAmountLend; // 结算时的实际出借金额
        uint256 settleAmountBorrow; // 结算时的实际借款金额
        uint256 finishAmountLend; //完成时的实际出借金额
        uint256 finishAmountBorrow; //完成时的实际借款金额, 还剩余多少borrowToken可以退给借款人
        uint256 liquidateAmountLend; //清算时的实际出借金额
        uint256 liquidateAmountBorrow; //清算时的实际借款金额
    }

    PoolDataInfo[] public poolDataInfo;

    // 借款人
    struct BorrowInfo {
        uint256 stakeAmount; // 当前借款的质押金额
        uint256 refundAmount; //多余退款金额
        bool hasRefund; // 默认false 未退款, true 已退款
        bool hasClaim; // 默认false 未领取, true 已领取
    }
    // Info of each user that stakes tokens.  {user.address : {pool.index : user.borrowInfo}}
    mapping(address => mapping(uint256 => BorrowInfo)) public userBorrowInfo;

    // 出借人
    struct LendInfo {
        uint256 stakeAmount; // 当前出借的质押金额
        uint256 refundAmount; //多余退款金额
        bool hasRefund; // 默认false 未退款, true 已退款
        bool hasClaim; // 默认false 未领取, true 已领取
        bool hasEmergencyWithdrawal; // 默认false 未紧急赎回, true 已紧急赎回
    }
    // Info of each user that stakes tokens.  {user.address : {pool.index : user.lendInfo}}
    mapping(address => mapping(uint256 => LendInfo)) public userLendInfo;

    // 有多少事件,就有多少方法
    // 存款事件, from借出者地址, token 借出者代币地址, amount 借出金额, mintAmount 生成数量
    event DepositLend(
        address indexed from,
        address indexed token,
        uint256 amount,
        uint256 mintAmount
    );
    // 借出退款事件, from退款者地址, token 退款者代币地址, refund 退款数量
    event RefundLend(
        address indexed from,
        address indexed token,
        uint256 refund
    );
    // 借出索赔事件, from索赔者地址, token 索赔者代币地址, refund 索赔数量
    event ClaimLend(
        address indexed from,
        address indexed token,
        uint256 refund
    );
    // 提取借出事件, from提取者地址, token 提取者代币地址, amount 提取金额, burnAmount 销毁数量
    event WithdrawLend(
        address indexed from,
        address indexed token,
        uint256 amount,
        uint256 burnAmount
    );
    // 存款借入事件, from借入者地址, token 借入的代币地址, amount 借入数量, mintAmount 生成数量
    event DepositBorrow(
        address indexed from,
        address indexed token,
        uint256 amount,
        uint256 mintAmount
    );
    // 借入退款事件, from退款者地址, token 退款者代币地址, refund 退款数量
    event RefundBorrow(
        address indexed from,
        address indexed token,
        uint256 refund
    );
    // 借入索赔事件, from索赔者地址, token 索赔者代币地址, refund 索赔数量
    event ClaimBorrow(
        address indexed from,
        address indexed token,
        uint256 amount
    );
    //  提取借入事件, from提取者地址, token 提取者代币地址, amount 提取金额, burnAmount 销毁数量
    event WithdrawBorrow(
        address indexed from,
        address indexed token,
        uint256 amount,
        uint256 burnAmount
    );
    // 交换事件, fromCoin 兑换前代币地址, toCoin 兑换后代币地址, fromValue 兑换前代币数量, toValue 兑换后代币数量
    event Swap(
        address indexed fromCoin,
        address indexed toCoin,
        uint256 fromValue,
        uint256 toValue
    );
    // 紧急借入提取事件, from提取者地址, token 提取者代币地址, amount 提取数量
    event EmergencyBorrowWithdrawal(
        address indexed from,
        address indexed token,
        uint256 amount
    );
    // 紧急借出提取事件, from提取者地址, token 提取者代币地址, amount 提取数量
    event EmergencyLendWithdrawal(
        address indexed from,
        address indexed token,
        uint256 amount
    );
    // 状态改变, pid 项目id
    event StateChange(
        uint256 indexed pid,
        PoolState indexed beforeState,
        PoolState indexed afterState
    );
    // 设置费用事件
    event SetFee(uint256 indexed newLendFee, uint256 indexed newBorrowFee);
    // 设置交换路由器地址事件
    event SetSwapRouterAddress(
        address indexed oldSwapAddress,
        address indexed newSwapAddress
    );
    // 设置费用地址事件
    event SetFeeAddress(
        address indexed oldFeeAddress,
        address indexed newFeeAddress
    );
    // 设置最小数量事件
    event SetMinAmount(
        uint256 indexed oldMinAmount,
        uint256 indexed newMinAmount
    );

    // modifier
    modifier notPause() {
        require(globalPause == false, "PledgePool: paused");
        _;
    }
    // 在结算时间之前
    modifier timeBefore(uint256 _pid) {
        require(
            block.timestamp < poolBaseInfo[_pid].settleTime,
            "PledgePool: time is over"
        );
        _;
    }

    modifier timeAfter(uint256 _pid) {
        require(
            block.timestamp > poolBaseInfo[_pid].settleTime,
            "PledgePool: time is not over"
        );
        _;
    }

    // 撮合中
    modifier stateMatch(uint256 _pid) {
        require(
            poolBaseInfo[_pid].state == PoolState.MATCH,
            "PledgePool: state is not settle"
        );
        _;
    }

     // 撮合中
    modifier stateExecution(uint256 _pid) {
        require(
            poolBaseInfo[_pid].state == PoolState.EXECUTION,
            "PledgePool: state is not execution"
        );
        _;
    }


    //  MATCH,
    //     EXECUTION,
    //     FINISH,
    //     LIQUIDATION,
    //     UNDONE
    // 撮合完成 未结束
    modifier stateNotMacthUndo(uint256 _pid) {
        PoolState state = poolBaseInfo[_pid].state;
        require(
            state == PoolState.EXECUTION ||
                state == PoolState.FINISH ||
                state == PoolState.LIQUIDATION,
            "PledgePool: state is not match or undo"
        );
        _;
    }

    modifier stateFinishLiquidation(uint256 _pid) {
        PoolState state = poolBaseInfo[_pid].state;
        require(
            state == PoolState.FINISH || state == PoolState.LIQUIDATION,
            "PledgePool: state is not liquidation"
        );
        _;
    }
    modifier stateUndone(uint256 _pid) {
        require(
            poolBaseInfo[_pid].state == PoolState.UNDONE,
            "PledgePool: state is not undone"
        );
        _;
    }

    // 获取最新的预言机价格
    function getUnderlyingPriceView(uint256 _pid) public view returns (uint256[2] memory) {
      PoolBaseInfo storage pool = poolBaseInfo[_pid];
      uint256[] memory assets = new uint256[](2);
      assets[0] = uint256(uint160(pool.lendToken));
      assets[1] = uint256(uint160(pool.borrowToken));
      uint256[] memory prices = oracle.getPrices(assets);
      return [prices[0], prices[1]];
    }

    constructor(
        address _oracle,
        address _swapRouter,
        address payable _feeAddress,
        address _multiSignature
    ) multiSignatureClient(_multiSignature) {
        require(
            _oracle != address(0),
            "PledgePool: oracle is the zero address"
        );
        require(
            _swapRouter != address(0),
            "PledgePool: swapRouter is the zero address"
        );
        require(
            _feeAddress != address(0),
            "PledgePool: feeAddress is the zero address"
        );

        oracle = IBscPledgeOracle(_oracle);
        swapRouter = _swapRouter;
        feeAddress = _feeAddress;
        lendFee = 0;
        borrowFee = 0;
    }

    function setFee(uint256 _lendFee, uint256 _borrowFee) external validCall {
        lendFee = _lendFee;
        borrowFee = _borrowFee;
        emit SetFee(_lendFee, _borrowFee);
    }

    function setSwapRouterAddress(address _swapRouter) external validCall {
        require(
            _swapRouter != address(0),
            "PledgePool: swapRouter is the zero address"
        );
        // 先emit, 再赋值. 如果先赋值,那么swapRouter就会变成新值,无法保留原有值
        emit SetSwapRouterAddress(swapRouter, _swapRouter);
        swapRouter = _swapRouter;
    }

    function setFeeAddress(address payable _feeAddress) external validCall {
        require(
            _feeAddress != address(0),
            "PledgePool: feeAddress is the zero address"
        );
        emit SetFeeAddress(feeAddress, _feeAddress);
        feeAddress = _feeAddress;
    }

    function setMinAmount(uint256 _minAmount) external validCall {
        emit SetMinAmount(minAmount, _minAmount);
        minAmount = _minAmount;
    }

    function poolLength() external view returns (uint256) {
        return poolBaseInfo.length;
    }

    function createPoolInfo(
        uint256 _settleTime,
        uint256 _endTime,
        uint256 _interestRate,
        uint256 _maxSupply,
        uint256 _mortgageRate,
        address _lendToken,
        address _borrowToken,
        address _spCoin,
        address _jpCoin,
        uint256 _autoLiquidateThreshold
    ) public validCall {
        require(
            _settleTime < _endTime,
            "PledgePool: settleTime must be less than endTime"
        );
        require(
            _jpCoin != address(0),
            "PledgePool: jpCoin is the zero address"
        );
        require(
            _spCoin != address(0),
            "PledgePool: spCoin is the zero address"
        );

        poolBaseInfo.push(
            PoolBaseInfo({
                settleTime: _settleTime,
                endTime: _endTime,
                interestRate: _interestRate,
                maxSupply: _maxSupply,
                lendSupply: 0,
                borrowSupply: 0,
                mortgageRate: _mortgageRate,
                lendToken: _lendToken,
                borrowToken: _borrowToken,
                state: defaultChoice,
                spCoin: IDebtToken(_spCoin),
                jpCoin: IDebtToken(_jpCoin),
                autoLiquidateThreshold: _autoLiquidateThreshold
            })
        );
        poolDataInfo.push(
            PoolDataInfo({
                settleAmountLend: 0,
                settleAmountBorrow: 0,
                finishAmountLend: 0,
                finishAmountBorrow: 0,
                liquidateAmountLend: 0,
                liquidateAmountBorrow: 0
            })
        );
    }

    function getPoolState(uint256 _pid) external view returns (uint256) {
        PoolBaseInfo storage pool = poolBaseInfo[_pid];
        return uint256(pool.state);
    }

    //  存款
    // _stakeAmount 质押数量
    function depositLend(
        uint256 _pid,
        uint256 _stakeAmount
    ) external payable nonReentrant notPause timeBefore(_pid) stateMatch(_pid) {
        PoolBaseInfo storage pool = poolBaseInfo[_pid];
        LendInfo storage lendInfo = userLendInfo[msg.sender][_pid];
        if (pool.lendToken == address(0)) {
            require(
                msg.value == _stakeAmount,
                "PledgePool: stakeAmount less than minAmount"
            );
            require(
                msg.value <= pool.maxSupply.sub(pool.lendSupply),
                "PledgePool: stakeAmount less than minAmount"
            );
        } else {
            require(
                _stakeAmount <= pool.maxSupply.sub(pool.lendSupply),
                "PledgePool: exceed maxSupply"
            );
        }
        uint256 amount = getPayableAmount(pool.lendToken, _stakeAmount);
        require(
            amount >= minAmount,
            "PledgePool: stakeAmount less than minAmount"
        );
        // 保存借款用户信息
        lendInfo.hasClaim = false;
        lendInfo.hasRefund = false;
        if (pool.lendToken == address(0)) {
            lendInfo.stakeAmount = lendInfo.stakeAmount.add(msg.value);
            pool.lendSupply = pool.lendSupply.add(msg.value);
        } else {
            lendInfo.stakeAmount = lendInfo.stakeAmount.add(_stakeAmount);
            pool.lendSupply = pool.lendSupply.add(_stakeAmount);
        }
        emit DepositLend(msg.sender, pool.lendToken, _stakeAmount, amount);
    }

    // 退还过量存款给存款人
    function refundLend(
        uint256 _pid
    ) external nonReentrant notPause timeAfter(_pid) stateNotMacthUndo(_pid) {
        PoolBaseInfo storage pool = poolBaseInfo[_pid];
        PoolDataInfo storage data = poolDataInfo[_pid];
        LendInfo storage lendInfo = userLendInfo[msg.sender][_pid];
        // 限制金额
        require(lendInfo.stakeAmount > 0, "PledgePool: stakeAmount is zero");
        // 实际存款数>实际借款数, 则需要退
        require(
            pool.lendSupply.sub(data.settleAmountLend) > 0,
            "refundLend: not refund"
        );
        // 未退款
        require(!lendInfo.hasRefund, "refundLend: has refund");

        // 用户持股比例
        uint256 userShare = lendInfo.stakeAmount.mul(calDecimal).div(
            pool.lendSupply
        );
        // 退款金额 = 总需退款金额 * 持股比例
        uint256 refundAmount = (pool.lendSupply.sub(data.settleAmountLend))
            .mul(userShare)
            .div(calDecimal);
        // 退款操作
        _redeem(payable(msg.sender), pool.lendToken, refundAmount);
        // 更新状态
        lendInfo.hasRefund = true;
        lendInfo.refundAmount = refundAmount;

        emit RefundLend(msg.sender, pool.lendToken, refundAmount);
    }

    // 存款人接收sp_token, 借出凭证
    function claimLend(
        uint256 _pid
    ) external nonReentrant notPause timeAfter(_pid) stateNotMacthUndo(_pid) {
        PoolBaseInfo storage pool = poolBaseInfo[_pid];
        PoolDataInfo storage data = poolDataInfo[_pid];
        LendInfo storage lendInfo = userLendInfo[msg.sender][_pid];
        // 限制金额
        require(lendInfo.stakeAmount > 0, "claimLend: stakeAmount is zero");
        require(lendInfo.hasClaim == false, "claimLend: has claim");
        // 计算份额 = 该用户质押金额 / 该池子总质押金额
        uint256 userShare = lendInfo.stakeAmount.mul(calDecimal).div(
            pool.lendSupply
        );
        // 计算可领取的sp_token数量
        uint256 claimAmount = data.settleAmountLend.mul(userShare).div(
            calDecimal
        );
        // 铸造sp_tiken
        pool.spCoin.mint(msg.sender, claimAmount);
        lendInfo.hasClaim = true;
        emit ClaimLend(msg.sender, pool.lendToken, claimAmount);
    }

    // 存款人取回本金和利息
    // 池子完成或被清算
    function withdrawLend(
        uint256 _pid,
        uint256 _spAmount
    )
        external
        nonReentrant
        notPause
        timeAfter(_pid)
        stateFinishLiquidation(_pid)
    {
        PoolBaseInfo storage pool = poolBaseInfo[_pid];
        PoolDataInfo storage data = poolDataInfo[_pid];
        LendInfo storage lendInfo = userLendInfo[msg.sender][_pid];
        // 限制金额
        require(_spAmount > 0, "withdrawLend: _spAmount is zero");
        // 销毁sp_token
        pool.spCoin.burn(msg.sender, _spAmount);
        // 计算销毁份额 (比例)
        uint256 shShare = _spAmount.mul(calDecimal).div(data.settleAmountLend);

        // 完成
        if (pool.state == PoolState.FINISH) {
            require(
                block.timestamp >= pool.endTime,
                "withdrawLend: not finish"
            );
            // 赎回金额
            uint256 redeemAmount = data.finishAmountLend.mul(shShare).div(
                calDecimal
            );
            _redeem(payable(msg.sender), pool.lendToken, redeemAmount);

            emit WithdrawLend(
                msg.sender,
                pool.lendToken,
                redeemAmount,
                _spAmount
            );
        }

        // 清算
        if (pool.state == PoolState.LIQUIDATION) {
            require(
                block.timestamp >= pool.settleTime,
                "withdrawLend: not after macth"
            );
            // 赎回金额
            uint256 redeemAmount = data.liquidateAmountLend.mul(shShare).div(
                calDecimal
            );
            _redeem(payable(msg.sender), pool.lendToken, redeemAmount);
            emit WithdrawLend(
                msg.sender,
                pool.lendToken,
                redeemAmount,
                _spAmount
            );
        }
    }

    //紧急提取
    function emergencyLendWithdrawal(
        uint256 _pid
    )
        external
        nonReentrant
        notPause
        timeAfter(_pid)
        stateFinishLiquidation(_pid)
    {
      PoolBaseInfo storage pool = poolBaseInfo[_pid];
      require(pool.lendSupply > 0,  "emergencyLendWithdrawal: lendSupply is zero");
      LendInfo storage lendInfo = userLendInfo[msg.sender][_pid];
      require(lendInfo.stakeAmount > 0, "emergencyLendWithdrawal: stakeAmount is zero");
      require(!lendInfo.hasEmergencyWithdrawal, "emergencyLendWithdrawal: has refund");
      // 赎回
      _redeem(payable(msg.sender), pool.lendToken, lendInfo.stakeAmount);
      lendInfo.hasEmergencyWithdrawal = true;
      emit EmergencyLendWithdrawal(msg.sender, pool.lendToken, lendInfo.stakeAmount);
    }

   
    // 是否能执行结算
    function checkoutSettle(uint256 _pid) public view returns (bool) {
      return block.timestamp >= poolBaseInfo[_pid].settleTime;
    }

    // 结算
    function settle(uint256 _pid) public validCall stateMatch(_pid) {
      PoolBaseInfo storage pool = poolBaseInfo[_pid];
      PoolDataInfo storage data = poolDataInfo[_pid];
      require(checkoutSettle(_pid), "settle: not settle time");

      if (pool.lendSupply > 0 && pool.borrowSupply > 0) {
        uint256[2] memory prices = getUnderlyingPriceView(_pid);
        // 总保证金价值 = 保证金数量 * 保证金价格
        uint256 totalValue = pool.borrowSupply.mul(prices[1]).div(calDecimal);
        // 转换成稳定币价值
        uint256 actualValue = totalValue.mul(pool.mortgageRate).div(baseDecimal);

        if (pool.lendSupply > actualValue) {
          // 总借款>总借出
          data.settleAmountLend = actualValue;
          data.settleAmountBorrow = pool.borrowSupply;
        } else {
          // 总借款<总借出
          data.settleAmountLend = pool.lendSupply;
          data.settleAmountBorrow = pool.lendSupply.mul(pool.mortgageRate).div(prices[1]).mul(baseDecimal).div(prices[0]);
        }
        pool.state = PoolState.EXECUTION;
        emit StateChange(_pid, PoolState.MATCH, PoolState.EXECUTION);
      } else {
        // 极端情况, 没有存款或借款
        pool.state = PoolState.UNDONE;
        data.settleAmountLend = pool.lendSupply;
        data.settleAmountBorrow = pool.borrowSupply;
        emit StateChange(_pid, PoolState.MATCH, PoolState.UNDONE);
      }
    }
    
    // 是否能finish
    function checkoutFinish(uint256 _pid) public view returns (bool) {
      return block.timestamp >= poolBaseInfo[_pid].endTime;
    }
    // 抵押借贷 + 到期自动结算 ,自动卖掉借款人抵押的 borrowToken，换回 lendToken，归还给出借人
    // 借款人在 settle后可以提现,此时finish时,和借款人关系不大,主要是此合约自动用借款人的borrowToke, 通过uniswap,兑换需要的lendToken, 付给出借人本金+利息和平台费用, 如果有多,则退还给出借人borrowToken
    // 借款人需要borrowToken, 出借人出借lendToken
    function finish(uint256 _pid) public validCall stateExecution(_pid) {
      PoolBaseInfo storage pool = poolBaseInfo[_pid];
      PoolDataInfo storage data = poolDataInfo[_pid];

      require(checkoutFinish(_pid), "finish: not finish time");

      // 计算比率 = (结束时间-结算时间)*基础小数 / 365天
      uint256 timeRatio = (pool.endTime.sub(pool.settleTime)).mul(baseDecimal).div(baseYear);

      // 计算利息 = 计算比率 * 年利率 * 结算贷款金额 注意这里年利率是用 1e16 精度（也就是 100% = 1e18）
      uint256 interest = timeRatio.mul(pool.interestRate.mul(data.settleAmountLend)).div(1e16);
      // 出借人应得金额: 计算贷款金额 = 结算贷款金额 + 利息
      uint256 lendAmount = data.settleAmountLend.add(interest);

      // 借款人到期后获得lendAmount, 根据lendAmount+平台费,计算借款人需要质押的borrowToken币数量
      // 借款人需要卖掉的borrowToken包含平台费,目标兑换量:计算销售金额 = 贷款金额  * (1+贷款费) 需要兑换的总数
      uint256 sellAmount = lendAmount.mul(lendFee.add(baseDecimal)).div(baseDecimal);

      // 根据出借人lendToken 和 借款人borrowToken, 及借款人需要卖掉的数量, 通过swap兑换 借款人卖掉borrowToken -> 获得lendToken,付给出借人和平台费
      // 根据需要兑换的数量sellAmount, 获取实际卖掉的数量amountSell 和 转入合约的lendToken数量amountIn, 用于支付出借人和平台费
      // amountSell 实际卖掉借款人质押的borrowToken数量, amountIn 兑换获得的lendToken
      (uint256 amountSell, uint256 amountIn) = _sellExactAmount(swapRouter, pool.borrowToken, pool.lendToken, sellAmount);

      // 兑换得来的数量必须大于等于出借人应得金额, 滑点风险
      require(amountIn >= lendAmount, "finish: amountIn < lendAmount");

      if (amountIn > lendAmount) {
        uint256 feeAmount = amountIn.sub(lendAmount);
        // 剩余的就是平台费用
        _redeem(payable(feeAddress), pool.lendToken, feeAmount);
        // 考虑滑点等因素
        data.finishAmountLend = amountIn.sub(feeAmount);
      } else {
        data.finishAmountLend = amountIn;
      }

      // 计算剩余的借款金额并赎回借款费
      // 借款人剩余的质押borrowToken = 借款人实际质押borrowToken: settleAmountBorrow - finish时实际卖出的borrowToken
      uint256 remainNowAmount = data.settleAmountBorrow.sub(amountSell);
      // 先把借款人退回费用结一下,剩余就是借款人可以退回的borrowToken
      uint256 remainBorrowAmount = redeemFees(borrowFee, pool.borrowToken, remainNowAmount);
      data.finishAmountBorrow = remainBorrowAmount;

      pool.state = PoolState.FINISH;

      emit StateChange(_pid, PoolState.EXECUTION, PoolState.FINISH);
    }

    // 费用计算并赎回
    function redeemFees(uint256 feeRatio, address token, uint256 amount) internal returns (uint256) {
      uint256 fee = amount.mul(feeRatio)/baseDecimal;
      if (fee > 0) {
        _redeem(payable(feeAddress), token, fee);
      }
      return amount.sub(fee);
    }

    // 检查是否可以清算
    function checkoutLiquidate(uint256 _pid) external view returns (bool) {
      PoolBaseInfo storage pool = poolBaseInfo[_pid];
      PoolDataInfo storage data = poolDataInfo[_pid];

      uint256[2] memory prices = getUnderlyingPriceView(_pid);
      // 借款人质押当前价值 borrowToken 折算成 lendToken 的价值
      uint256 borrowValueNow = data.settleAmountBorrow.mul(prices[1].mul(calDecimal).div(prices[0])).div(calDecimal);
      // 清算阈值
      uint256 valueThreshold = data.settleAmountLend.mul(baseDecimal.add(pool.autoLiquidateThreshold)).div(baseDecimal);
      // 超额抵押 settleAmountLend实际出借数量即需要还的lendToken数量, 超额autoLiquidateThreshold(20%)抵押, 即阈值为settleAmountLend*(1+20%)
      return borrowValueNow < valueThreshold;
    }

    // 清算
    function liquidate(uint256 _pid) public validCall stateExecution(_pid) {
      PoolBaseInfo storage pool = poolBaseInfo[_pid];
      PoolDataInfo storage data = poolDataInfo[_pid];

      // 检查时间
      require(block.timestamp >= pool.settleTime, "liquidate: not settle");
      // 计算时间比率
      uint256 timeRatio = (pool.endTime.sub(pool.settleTime)).mul(baseDecimal).div(baseYear);
      uint256 interest = timeRatio.mul(pool.interestRate.mul(data.settleAmountLend)).div(1e16);
      uint256 lendAmount = data.settleAmountLend.add(interest);

      uint256 sellAmount = lendAmount.mul(lendFee.add(baseDecimal)).div(baseDecimal);
      (uint256 amountSell, uint256 amountIn) = _sellExactAmount(swapRouter, pool.borrowToken, pool.lendToken, sellAmount);

      // 清算及时止损,即使兑换的lendToken数量小于应得金额, 也要清算
      // require(amountIn >= lendAmount, "finish: amountIn < lendAmount");

      if (amountIn > lendAmount) {
        uint256 feeAmount = amountIn.sub(lendAmount);
        _redeem(payable(feeAddress), pool.lendToken, feeAmount);
        data.liquidateAmountLend = amountIn.sub(feeAmount);
      } else {
        data.liquidateAmountLend = amountIn;
      }

      uint256 remainNowAmount = data.settleAmountBorrow.sub(amountSell);
      uint256 remainBorrowAmount = redeemFees(borrowFee, pool.borrowToken, remainNowAmount);
      data.liquidateAmountBorrow = remainBorrowAmount;
      pool.state = PoolState.LIQUIDATION;
      emit StateChange(_pid, PoolState.EXECUTION, PoolState.LIQUIDATION);
    }

    //
    function _getSwapPath(address _swapRouter, address token0, address token1) internal pure returns (address[] memory path) {
      IUniswapV2Router02 IUniswap = IUniswapV2Router02(_swapRouter);
      path = new address[](2);
      path[0] = token0 == address(0) ? IUniswap.WETH() : token0;
      path[1] = token1 == address(0) ? IUniswap.WETH() : token1;
    }

    function _getAmountIn(address _swapRouter, address token0, address token1, uint256 amountOut) internal view returns (uint256) {
      IUniswapV2Router02 IUniswap = IUniswapV2Router02(_swapRouter);
      address[] memory path = _getSwapPath(_swapRouter, token0, token1);
      uint256[] memory amounts = IUniswap.getAmountsIn(amountOut, path);
      return amounts[0];
    }

    function _sellExactAmount(address _swapRouter, address token0, address token1, uint256 amountOut) internal returns (uint256, uint256) {
      uint256 amountSell = amountOut > 0?_getAmountIn(_swapRouter, token0, token1, amountOut):0;
      return (amountSell, _swap(_swapRouter, token0, token1, amountSell));
    }

    function _swap(address _swapRouter, address token0, address token1, uint256 amount) internal returns (uint256) {
      if (token0 != address(0)) {
        _safeApprove(token0, address(_swapRouter), type(uint256).max);
      }
      if (token1 != address(0)) {
        _safeApprove(token1, address(_swapRouter), type(uint256).max);
      }
      IUniswapV2Router02 IUniswap = IUniswapV2Router02(_swapRouter);
      address[] memory path = _getSwapPath(_swapRouter, token0, token1);
      uint256[] memory amounts;
      if (token0 == address(0)) {
        amounts = IUniswap.swapExactETHForTokens{value: amount}(0, path, address(this), block.timestamp+30);
      } else if (token1 == address(0)) {
        amounts = IUniswap.swapExactTokensForETH(amount, 0, path, address(this), block.timestamp+30);
      } else {
        amounts = IUniswap.swapExactTokensForTokens(amount, 0, path, address(this), block.timestamp+30);
      }
      emit Swap(token0, token1,amounts[0], amounts[amounts.length-1]);
      return amounts[amounts.length-1];
    }

    function _safeApprove(address token, address to, uint256 value) internal {
      (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0x095ea7b3, to, value));
      require(success && (data.length == 0 || abi.decode(data, (bool))), 'TransferHelper: APPROVE_FAILED');
    }

}
