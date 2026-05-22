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
./bc send -from '[\"Kunyan\", \"Kunyan\", \"CZ\"]' -to '[\"WU\", \"CZ\", \"Kunyan\"]' -amount '[\"2\", \"3\", \"1\"]'

./bc getbalance -address "WU"

# 打印所有区块
./bc printchain

# 测试完后记得删除 db
Remove-Item -Recurse -Force blockchain.db, bc.exe
```


## 用 AI vibe coding 把PoW改成别的共识算法（如 DAG-based， DPOS，PBFT...)