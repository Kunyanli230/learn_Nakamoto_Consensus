package BLC

func (cli *CLI) createGenesisBlockchain(address string, nodeID string) {
	blockchain := CreateBlockchainWithGenesisBlock(address, nodeID)
	defer blockchain.DB.Close()

	utxoSet := &UTXOSet{blockchain}
	utxoSet.ResetUTXOSet()
}
