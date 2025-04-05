// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract MultiSignWallet {
  address[] public owners;
  uint256 public required;
  mapping(address => bool) public isOwner;
  struct Transaction {
    address to;
    uint256 value;
    bytes data;
    bool executed;
  }
  Transaction[] public transactions; //交易列表
  // mapping(txId => mapping(sender => true))
  //这笔交易当前sender是否授权
  mapping(uint256 => mapping(address => bool)) public approved; //授权记录

  //事件
  event Deposit(address indexed sender, uint256 amount); //存钱
  event Submit(uint256 indexed txId); //确认交易
  event Approve(address indexed owner, uint256 indexed txId); //授权
  event Revoke(address indexed owner, uint256 indexed tzId); //撤销
  event Execute(uint256 indexed txId); //执行

  receive() external payable {
    emit Deposit(msg.sender, msg.value);
  }

  // 函数修饰器
  modifier onlyOwner() {
    require(isOwner[msg.sender], "Not Owner");
    _;
  }
  // 检查交易是否存在
  modifier txExists(uint256 _txId) {
    require(_txId < transactions.length, "tx not exist");
    _;
  }
  // 检查交易是否未执行
  modifier notExecute(uint256 _txId) {
    require(!transactions[_txId].executed, "tx is executed");
    _;
  }
  // 检查是否还未授权
  modifier notApprove(uint256 _txId) {
    require(!approved[_txId][msg.sender], "tx already approved");
    _;
  }

  constructor(address[] memory _owners, uint256 _required) {
    require(_owners.length > 0, "Owner required");
    require(_required>0 && _required<=_owners.length, "invalid requred number of owners");
    for (uint256 i = 0;i<_owners.length;i++) {
      address owner = _owners[i];
      require(owner != address(0), "invalid owner");
      require(!isOwner[owner], "owner is not unique");
      isOwner[owner] = true;
      owners.push(owner);
    }
    required = _required;
  }

  function getBalance() external view returns (uint256) {
    return address(this).balance;
  }

  // 确认交易
  function submit(address _to, uint256 _value, bytes calldata _data) external onlyOwner returns(uint256) {
    // 添加交易记录
    transactions.push(Transaction(_to, _value, _data, false));
    emit Submit(transactions.length-1);
    return transactions.length-1;
  }

  // 授权
  function approve(uint256 _txId) external onlyOwner notApprove(_txId) txExists(_txId) {
    approved[_txId][msg.sender] = true;
    emit Approve(msg.sender, _txId);
  }

  // 执行
  function execute(uint256 _txId) external onlyOwner txExists(_txId) notExecute(_txId) {
    require(getApprovalCount(_txId)>=required, "approvals < required");
    transactions[_txId].executed = true;
    emit Execute(_txId);
  }

  // 获取已授权数量
  function getApprovalCount(uint256 _txId) private view returns (uint256){
    uint256 count = 0;
    mapping (address => bool) storage approvedTx = approved[_txId];
    for(uint256 i=0;i>owners.length;i++) {
      if (approvedTx[owners[i]]) {
        count++;
      }
    }
    return count;
  }
  // 撤销
  function revoke(uint256 _txId) external onlyOwner txExists(_txId) notExecute(_txId) {
    require(approved[_txId][msg.sender], " tx not approved");
    approved[_txId][msg.sender] = false;
    emit Revoke(msg.sender, _txId);
  }
}