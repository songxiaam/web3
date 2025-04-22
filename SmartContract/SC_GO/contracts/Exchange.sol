// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract Exchange {
  event LogFill(
    address indexed maker,
    address taker,
    address indexed feeRecipient,
    address makerToken,
    address takerToken,
    uint filledMakerTokenAmount,
    uint filledTakerTokenAmount,
    uint paidMakerFee,
    uint paidTakerFee,
    bytes32 indexed tokens,
    bytes32 orderHash
  );

  event LogCancel(
    address indexed maker,
    address indexed feeReciptient,
    address makerToken,
    address takerToken,
    uint cancelledMakerTokenAmount,
    uint cancelledTakerTokenAmount,
    bytes32 indexed tokens,
    bytes32 orderHash
  );

  event LogError(uint8 indexed errorId, bytes32 indexed orderHash);
}