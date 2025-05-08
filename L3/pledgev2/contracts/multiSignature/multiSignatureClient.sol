// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

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