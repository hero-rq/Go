package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data    string
	hash    string
	preHash string
}

type blockchain struct {
	blocks []block
}

func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

func (b *blockchain) getAdd(data string) {
	newBlock := block{data, "", b.getLastHash()}
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.preHash))
	hexHash := fmt.Sprintf("%x", hash)
	newBlock.hash = hexHash
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) getList() {
	for _, block := range b.blocks {
		fmt.Printf("Data : %s\n", block.data)
		fmt.Printf("Hash : %s\n", block.hash)
		fmt.Printf("PrevHash : %s\n", block.preHash)
	}
}

func main() {
	chain := blockchain{}
	chain.getAdd("GenesisBlock")
	chain.getAdd("SecondBlock")
	chain.getAdd("ThirdBlock")
	chain.getList()
}
