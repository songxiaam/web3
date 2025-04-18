
纯函数（pure function）是函数式编程中的一个重要概念。在编程中，如果一个函数满足以下条件，则它可以被认为是纯函数：

1. 相同输入产生相同输出：对于相同的输入值，纯函数必须始终产生相同的输出。
2. 输出仅依赖输入：函数的输出只依赖于其输入参数，而与其他隐藏信息或状态无关，也不受 I/O 设备产生的外部输出影响。
3. 无副作用：函数不能有语义上可观察的副作用，比如触发事件、输出到设备、或更改输入参数以外的内容等。

简单来说，纯函数不读不写状态变量，没有副作用。使用纯函数可以提高代码的安全性，避免出现与预期不符的副作用。

如果你的函数承诺不需要查询，也不需要修改合约状态，那么你应该为它加上 `pure` 修饰符。如下所示：

`pure` 函数

```
function add(uint lhs, uint rhs) public pure returns(uint) {
    return lhs + rhs;
}
```

被标记为 `pure` 的函数，如果你在函数体查询或者修改合约状态，编译器都会直接报错。

## 怎样才算查询合约状态

Solidity 有 5 种行为被认为是查询了合约状态：

1. 读取状态变量：访问和读取合约中的状态变量。
2. 访问 `address(this).balance` 或 `<address>.balance`：获取当前合约或指定地址的余额。
3. 访问 `block`、`tx`、`msg` 的成员：读取区块链相关信息，如 `block.timestamp`、`tx.gasprice`、`msg.sender` 等。
4. 调用未标记为 `pure` 的任何函数：调用任何未被标记为 `pure` 的函数，即使这些函数本身也只是读取状态。
5. 使用包含某些操作码的内联汇编：使用可能会读取状态的特定汇编代码。

## 怎样才算修改合约状态

Solidity 有 8 种行为被认为是修改了合约状态：

1. 修改状态变量
2. 触发事件
3. 创建其他合约
4. 使用 `selfdestruct` 来销毁合约
5. 通过函数调用发送以太币
6. 调用未标记为 `view` 或 `pure` 的任何函数
7. 使用低级别调用，如 `transfer`, `send`, `call`, `delegatecall` 等
8. 使用包含某些操作码的内联汇编
