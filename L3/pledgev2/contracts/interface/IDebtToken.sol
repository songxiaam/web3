// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

interface IDebtToken {

  function balanceOf(address account) external view returns (uint256);
  function totalSupply() external view returns (uint256);
  // 铸造/新增, 创建新的代币
  function mint(address account, uint256 amount) external;
  // 销毁
  function burn(address account, uint256 amount) external;
}