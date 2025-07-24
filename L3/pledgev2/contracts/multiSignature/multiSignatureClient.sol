// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

/**
这是一个多签名合约的客户端基类，主要用于为其他合约提供多签名校验功能。
通过修饰器 validCall，可以让合约的关键操作必须经过多签名合约的验证，提升安全性。
适用于需要多方共同管理和授权的场景。
 */

interface IMultiSignature {
  function getValidSignature(bytes32 msghash, uint256 lastIndex) external view returns (uint256);
}

contract multiSignatureClient {
  uint256 private constant multiSignaturePosition = uint256(keccak256("org.multiSignature.storage"));
  uint256 private constant defaultIndex = 0;

  constructor(address multiSignature) {
    require(multiSignature != address(0), "multiSignatureClient: multiSignature is the zero address");
    saveValue(multiSignaturePosition, uint256(uint160(multiSignature)));
  }

  function getMultiSignatureAddress() internal view returns (address) {
    return address(uint160(getValue(multiSignaturePosition)));
  }

  modifier validCall() {
    checkMultiSignature();
    _;
  }

  function checkMultiSignature() internal view {
    uint256 value;
    assembly {
      value := callvalue()
    }
    bytes32 msgHash = keccak256(abi.encodePacked(msg.sender, address(this)));
    address multiSign = getMultiSignatureAddress();
    uint256 newIndex = IMultiSignature(multiSign).getValidSignature(msgHash, value);
    require(newIndex > defaultIndex, "multiSignatureClient : This tx is not aprroved");
  }

  function saveValue(uint256 position, uint256 value) internal {
    assembly {
      sstore(position, value)
    }
  }

  function getValue(uint256 position) internal view returns (uint256 value) {
    assembly {
      value := sload(position)
    }
  }
}