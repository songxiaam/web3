// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract Error {
  function testRequire(uint256 _i) public pure {
    require (_i<10, "Input must be less than 10");
  }

  function testRevert(uint256 _i) public pure {
    if (_i<20) {
      revert("Input must be greater than 20");
    }
  }

  uint256 public num;
  function testAssert() public view {
    assert(num==0);
  }

  error InsufficientBalance(uint256 balance, uint256 withdrawAmount);

  function testCustomError(uint256 _withdrawAmount) public view {
    uint256 bal = address(this).balance;
    if (bal < _withdrawAmount) {
      revert InsufficientBalance({
        balance:bal,
        withdrawAmount: _withdrawAmount
      });
    }
  }
}