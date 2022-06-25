package main

import (
	"fmt"

	"github.com/hero-rq/learngo/blockchain"
)

func main() {
	chain := blockchain.GetblockChain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")
	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Prevhash: %s\n", block.PrevHash)
	}
}
