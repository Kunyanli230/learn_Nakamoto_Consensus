package BLC

import "fmt"

func (cli *CLI) CreateWallet(nodeID string) {
	wallets, _ := NewWallets(nodeID)
	wallets.CreateNewWallet(nodeID)
	fmt.Println(len(wallets.WalletsMap))
}
