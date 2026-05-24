# 公链项目

一步一步实现公链

## Background

建设中
* 建设中

## CLI用法
```
go build -o bc.exe main.go

# 创建创世区块
./bc createblockchain -address "Kunyan"

# 查询余额
./bc getbalance -address "Kunyan"

# 转账（PowerShell 需用 \" 转义内层双引号）
./bc send -from '[\" "]' -to '[\" "]' -amount '[\"2\"]'

./bc send -from '[\"1JjTSv7y4sqgSixUAV67cFcX4zGzJzz6EH\"]' -to '[\"1BPXKn98Xryng8YBirPwv6Gj79yjPduiK8\"]' -amount '[\"2\"]'


./bc getbalance -address "1BPXKn98Xryng8YBirPwv6Gj79yjPduiK8"

# 打印所有区块
./bc printchain

# 测试完后记得删除 db
Remove-Item -Recurse -Force blockchain.db, bc.exe, wallets.dat
```


