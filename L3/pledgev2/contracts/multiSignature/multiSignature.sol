// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "./multiSignatureClient.sol";

library whiteListAddress {
  function addWhiteListAddress(address[] storage whiteList, address temp) internal {
    if (!isEligibleAddress(whiteList, temp)) { 
      whiteList.push(temp);
    }
  }

  function removeWhiteListAddress(address[] storage whiteList, address temp) internal returns (bool) {
    uint256 len = whiteList.length;
    uint256 i= 0;
    for (;i<len;i++) {
      if (whiteList[i] == temp) {
        break;
      }
    }
    if (i < len) {
      if (i != len - 1) {
        whiteList[i] = whiteList[len - 1];
      }
      whiteList.pop();
      return true;
    }
    return false;
  }

  // 是否已包含
  function isEligibleAddress(address[] memory whiteList, address temp) internal pure returns (bool) {
    uint256 len = whiteList.length;
    for (uint256 i = 0; i < len; i++) {
      if (whiteList[i] == temp) {
        return true;
      }
    }
    return false;
  }
}