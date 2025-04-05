// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract Bank {
  address public immutable owner;

  event Deposit(address _ads, uint256 amount);
  event Withdraw(uint256 amount);

  receive() external payable {
    emit Deposit(msg.sender, msg.value);
  }

  constructor() payable {
    owner = msg.sender;
  }

  function withdraw() external {
    require (msg.sender == owner, "Not Owner");
    emit Withdraw(address(this).balance);
    selfdestruct(payable(msg.sender));
  }

  function getBalance() external view returns (uint256) {
    return address(this).balance;
  }

}