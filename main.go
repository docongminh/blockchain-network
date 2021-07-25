package main

import (
	"blockchain/core"
	"fmt"
)

func main() {
	bc := core.NewBlockChain()
	bc.AddBlock("Send 50 MCoin to me")
	bc.AddBlock("Send 100 Mcoin to you")
	for _, block := range bc.Blocks {
		fmt.Printf("Prev has: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("----------------------\n")
	}
}
