// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract Shipping {
  enum ShippingStatus {
    Pending, Shipped, Delivered
  }

  ShippingStatus private status;

  event LogNewAlert(string desc);

  constructor() {
    status = ShippingStatus.Pending;
  }

  function shipped() public {
    status = ShippingStatus.Shipped;
    emit LogNewAlert("Shipped");
  }

  function delivered() public {
    status = ShippingStatus.Delivered;
    emit LogNewAlert("Delivered");
  }

  function getStatus(ShippingStatus _status) public pure returns (string memory) {
    if (ShippingStatus.Pending == _status) return "Pending"; 
    if (ShippingStatus.Shipped == _status) return "Shipped"; 
    if (ShippingStatus.Delivered == _status) return "Delivered"; 
    return "UnKnouw";
  }

  function Status() public view returns (string memory) {
      ShippingStatus _status = status;
      return getStatus(_status);
  }

}