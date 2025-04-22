// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract MyErc20 {
  event Transfer(address indexed from, address indexed to, uint tokens);
  event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}