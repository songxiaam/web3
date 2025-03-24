// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract RCCToken is ERC20 {
    // 合约内容
    constructor() ERC20("RccToken", "RCC") {
      _mint(msg.sender, 20000000*1_000_000_000_000_000_000);
    }


}
