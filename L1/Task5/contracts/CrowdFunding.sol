// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract CrowdFunding {
  address public immutable beneficiary; //受益人
  uint256 public immutable fundingGoal; //目标金额
  uint256 public fundingAmount; //当前金额

  // 捐赠人 => 金额
  mapping(address => uint256) public funders;
  //
  mapping(address => bool) private fundersInserted;

  address[] public fundersKey; //length
  
  //
  bool public AVAILABLED = true;

  constructor(address _bebeficiary, uint256 _goal) {
    beneficiary = _bebeficiary;
    fundingGoal = _goal;
  }

  //资助
  function contribute() external payable {
    require(AVAILABLED, "CrowdFunding is closed");
    //检查捐赠金额是否超过目标,超过增退还
    uint256 potentialFundingAmount = fundingAmount+msg.value;
    uint256 refundAmount = 0;
    if(potentialFundingAmount > fundingGoal) {
      refundAmount = potentialFundingAmount-fundingGoal;
      funders[msg.sender] += msg.value-refundAmount;
      fundingAmount = fundingGoal;
    } else {
      funders[msg.sender] += msg.value;
      fundingAmount += msg.value;
    }

    // 更新捐助人列表
    if (!fundersInserted[msg.sender]) {
      fundersInserted[msg.sender] = true;
      fundersKey.push(msg.sender);
    }

    // 退还多余金额
    if (refundAmount > 0) {
      payable(msg.sender).transfer(refundAmount);
    }
  }

  // 关闭
  function close() external returns(bool) {
    // 1.检查
    if (fundingAmount < fundingGoal) {
      return false;
    }

    uint256 amount = fundingAmount;
    // 2.修改
    fundingAmount = 0;
    AVAILABLED = false;

    // 3.操作
    payable(beneficiary).transfer(amount);
    return true;
  }

  function fundersLength() public view returns (uint256) {
    return fundersKey.length;
  }

}