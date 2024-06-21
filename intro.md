# 创建模块, 依赖于account和bank
`ignite scaffold module tokenfactory --dep account,bank`



# 定义 Denom 数据结构
```
Denom 是 "Denomination" 的缩写，指的是代币的名称或标识符。
```

要管理你的代币工厂中的代币类型，需要使用 Ignite map来定义它们的结构。这样可以将数据存储为键-值对。运行这个命令：

`ignite scaffold map Denom description:string ticker:string precision:int url:string maxSupply:int supply:int canChangeMaxSupply:bool --signer owner --index denom --module tokenfactory`

我来解释一下命令 : 

* **`ignite scaffold map`**： 这部分命令用于使用 Ignite CLI 创建一个新的数据映射结构。
* **`Denom`**： 这是要创建的映射结构的名称，它将用来存储代币类型的信息。
* **`description:string ticker:string precision:int url:string maxSupply:int supply:int canChangeMaxSupply:bool`**： 这些是映射结构中的字段，它们定义了代币类型的属性，例如描述、符号、精度、URL、最大供应量、供应量和是否可更改最大供应量。
* **`--signer owner`**： 它指定了谁拥有创建和修改代币类型的权限。在这个例子中，`owner` 表示只有拥有 `owner` 角色的账户才能创建和修改代币类型。
* **`--index denom`**： 这表示 `denom` 字段将被用作索引，方便快速查找特定的代币类型。
* **`--module tokenfactory`**： 这表示这个映射结构属于 `tokenfactory` 模块。


# 搭建新消息

该消息允许创建（铸造）新代币并将其分配给指定的接收者。必要的输入包括面额、铸币数量和收件人地址。

```
ignite scaffold message MintAndSendTokens denom:string amount:int recipient:string --module tokenfactory --signer owner
```

该消息促进了面额所有权的转移。它需要面额名称和新所有者的地址。

```
ignite scaffold message UpdateOwner denom:string newOwner:string --module tokenfactory --signer owner
```





