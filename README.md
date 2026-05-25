# 公链项目

一步一步实现公链

## Background

建设中
* 建设中

## CLI用法
### 1. 构建 CLI

```powershell
go build -o bc.exe main.go
```

### 2. 查看支持的命令

```powershell
./bc.exe
```

### 3. 创建钱包

先创建至少两个钱包地址，一个作为创世区块奖励地址，一个作为收款地址。

```powershell
./bc.exe createwallet
./bc.exe createwallet
./bc.exe createwallet
```


```powershell
./bc.exe addresslists
```

### 4. 创建区块链

```powershell
./bc.exe createblockchain -address "<创世区块奖励地址>"
```

调试的时候可以用这里：
钱包一：
钱包二：
钱包三：

### 5. 查询余额

```powershell
./bc.exe getbalance -address "<钱包地址>"
```

### 6. 转账

```powershell
./bc.exe send -from '[\"钱包一\",\"钱包一\"]' -to '[\"钱包二\",\"钱包三\"]' -amount '[\"2\",\"3\"]'

```
### 7. 打印区块链

```powershell
./bc.exe printchain
```

### 8. 清理本地运行文件
```powershell
Remove-Item -Force blockchain.db, bc.exe, wallets.dat
```


