package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// UTXO交易模型
type Transaction struct {
	//1.交易hash ()
	TxHash []byte

	//2.输入
	Vins []*TXInput
	
	//3.输出
	Vouts []*TXOutput
}

//transaction创建分两种情况
// 创世区块创建时的transaction
func NewCoinbaseTransaction(address string) *Transaction {
   
	txInput := &TXInput{[]byte{}, -1,"Genesis Block"}
	
	txOuptut := &TXOutput{10, address}
    
	txCoinbase := &Transaction{[]byte{}, []*TXInput{txInput}, []*TXOutput{txOuptut}}

	//设置hash值
	txCoinbase.HashTransaction()

	return txCoinbase
}

func (tx *Transaction) HashTransaction() {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]
}

// 转账时产生的transaction