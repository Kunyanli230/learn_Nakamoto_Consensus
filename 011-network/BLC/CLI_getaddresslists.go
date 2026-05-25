package BLC

import "fmt"

func (cli *CLI) addressLists(nodeID string) {
	fmt.Println("所有钱包地址如下:")
	wallets, _ := NewWallets(nodeID)
	for address := range wallets.WalletsMap {
		fmt.Println(address)
	}
}
