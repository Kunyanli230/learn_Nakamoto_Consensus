package BLC

import "fmt"

// 转账
func (cli *CLI) send(from []string, to []string, amount []string, nodeID string, mineNow bool) {
	blockchain := BlockchainObject(nodeID)
	defer blockchain.DB.Close()

	if mineNow {
		blockchain.MineNewBlock(from, to, amount, nodeID)
		utxoSet := &UTXOSet{blockchain}
		utxoSet.Update()
	} else {
		fmt.Println("由矿工节点处理...")
	}
}
