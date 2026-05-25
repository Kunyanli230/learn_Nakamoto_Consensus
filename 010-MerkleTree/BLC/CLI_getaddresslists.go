package BLC

import "fmt"

func (cli *CLI) addressLists() {
	fmt.Println("所有钱包地址如下:")
	wallets, _ := NewWallets()
	for address, _ := range wallets.WalletsMap {
		fmt.Println(address)
	}
}