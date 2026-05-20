# 公链项目

一步一步实现公链

## Background

建设中
* 建设中

## CLI用法
```
go build -o bc.exe main.go
./bc createblockchain -data "随便写“
./bc addblock -data "随便写“
./bc printchain
./bc send -from '["Kunyan", "CZ"]' -to '["WU", "PP"]' -amount '["10000","30000"]'

//测试完后记得删除db
Remove-Item -Recurse -Force blockchain.db, bc.exe


```


## 用 AI vibe coding 把PoW改成别的共识算法（如 DAG-based， DPOS，PBFT...)