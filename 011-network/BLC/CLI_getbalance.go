package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) getBalance(address string, nodeID string) {
	if DBExists(nodeID) == false {
		fmt.Println("数据库不存在.....")
		os.Exit(1)
	}
	blockchain := BlockchainObject(nodeID)
	defer blockchain.DB.Close()

	fmt.Println("地址： ", address)
	amount := blockchain.GetBalance(address)
	fmt.Printf("%s 的余额一共有 %d 个Token\n", address, amount)
}
