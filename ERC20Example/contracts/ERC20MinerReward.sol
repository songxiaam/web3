// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract ERC20MinerReward is ERC20 {
  event LogNewAlert(string desc, address indexed _from, uint256 _n);

  constructor() ERC20("MinerReward", "MRW"){
  }

  function _reward() public {
    _mint(block.coinbase, 20);
    emit LogNewAlert("_rewarded", block.coinbase, block.number);
  }
  
}