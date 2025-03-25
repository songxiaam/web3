// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract TodoList {
  uint public taskCount = 0;
  struct Task {
    uint id;
    string taskname;
    TaskStatusEnum status;
  }

  enum TaskStatusEnum {
    Pending,
    Processing,
    Finished,
    Canceled
  }

  mapping (uint => Task) public tasks;

  event TaskCreate(uint id, string taskName, TaskStatusEnum status);
  event TaskStatus(uint id, TaskStatusEnum status);
  constructor() {
    createTask("Todo List Tutorial...");
  }

  function createTask(string memory _taskName) public {
    taskCount++;
    Task memory newTask = Task(taskCount, _taskName, TaskStatusEnum.Pending);
    tasks[taskCount] = newTask;
    emit TaskCreate(taskCount, _taskName, TaskStatusEnum.Pending);
  }
  function getStatus(uint _id) public {
    Task storage task = tasks[_id];
    emit TaskStatus(_id, task.status);
  }

  function handling(uint _id) public {
      Task storage task = tasks[_id];
      task.status = TaskStatusEnum.Processing;
      emit TaskStatus(_id, task.status);
  }

  function finish(uint _id) public {
      Task storage task = tasks[_id];
      task.status = TaskStatusEnum.Finished;
      emit TaskStatus(_id, task.status);
  }

  function cancel(uint _id) public {
      Task storage task = tasks[_id];
      task.status = TaskStatusEnum.Canceled;
      emit TaskStatus(_id, task.status);
  }
}