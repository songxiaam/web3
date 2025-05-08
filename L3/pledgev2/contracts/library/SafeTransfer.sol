// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "./SafeErc20.sol";
import "../interface/IERC20.sol";

contract SafeTransfer {
  using SafeERC20 for IERC20;

  // 赎回
  event Redeem(address indexed recieptor, address indexed token, uint256 amount);

  function getPayableAmount(address token, uint256 amount) internal returns (uint256) {
    if (token == address(0)) {
      amount = msg.value;
    } else if (amount > 0) {
      IERC20 oToken = IERC20(token);
      oToken.safeTransferFrom(msg.sender, address(this), amount);
    }
    return amount;
  }

  // 赎回
  function _redeem(address payable recieptor, address token, uint256 amount) internal {
    if (token == address(0)) {
      recieptor.transfer(amount);
    } else {
      IERC20 oToken = IERC20(token);
      oToken.safeTransfer(recieptor, amount);
    }
    emit Redeem(token, recieptor, amount);
  }

}