package main

import (
	"fmt"
	"publicChain/001-block/BLC"
)

func main() {
	block := BLC.NewBlock("Genesis Block", 1, make([]byte, 32))

	fmt.Println(block)
}
