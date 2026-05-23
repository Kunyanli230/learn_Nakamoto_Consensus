package BLC

import (
	"fmt"
	"os"
)

// 查询余额
func (cli *CLI) getBalance(address string) {
	if DBExists() == false {
		fmt.Println("数据库不存在.....")
		os.Exit(1)
	}
	blockchain := BlockchainObject()
	defer blockchain.DB.Close()

	fmt.Println("地址： ", address)
	amount := blockchain.GetBalance(address)
	fmt.Printf("%s 的余额一共有 %d 个Token\n", address, amount)

}
