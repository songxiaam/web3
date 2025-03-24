// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract FunctionModifier {
  address public owner;
  uint256 public x=10;
  bool public locked;

  constructor() {
    owner = msg.sender;
  }

  modifier onlyOwner() {
    require(msg.sender == owner, "Not owner");
    _;
  }
  modifier validAddress(address _addr) {
    require(_addr != address(0), "Not valid address");
    _;
  }
  function changeOwner(address _newOwner) public onlyOwner validAddress(_newOwner) {
    owner = _newOwner;
  }

  modifier onReentrancy() {
    require(!locked, "No reentrancy");
    locked = true;
    _;
    locked = false;
  }

  function decreament(uint256 i) public onReentrancy {
    x -= i;
    if (i>1) {
      decreament(i-1);
    }
  }
}