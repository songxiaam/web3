// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract ToDOList {
  struct  ToDo {
    string name;
    bool isCompleted;
  }

  ToDo[] public list;

  function create(string memory name) public {
    ToDo memory item = ToDo(name, false);
    list.push(item);
  }

  function modifierToDoName(uint256 _index, string memory _name) external {
    list[_index].name = _name;
  }

  function modifierToDoName2(uint256 _index, string memory _name) external {
    ToDo storage temp = list[_index];
    temp.name = _name;
  }

  function modifierStatus(uint256 _index, bool _status) external {
    list[_index].isCompleted = _status;
  }

  function modifierStatus2(uint256 _index, bool _status) external {
    ToDo storage temp = list[_index];
    temp.isCompleted = _status;
  }

  //自动切换状态
  function modifierStatus3(uint256 _index) external {
    ToDo storage temp = list[_index];
    temp.isCompleted = !temp.isCompleted;
  }

  //获取任务
  function get(uint256 _index) external view returns (string memory, bool) {
    ToDo storage temp = list[_index]; //1次拷贝,省gas
    // ToDo memory temp2 = list[_index]; //2次拷贝
    return (temp.name, temp.isCompleted);
  }



}