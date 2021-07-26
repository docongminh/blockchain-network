package main

import (
	"blockchain/core"
	"fmt"
)

func main() {
	bc := core.NewBlockChain()
	bc.AddBlock("You got 50 MCoin")
	bc.AddBlock("You 25 Mcoin")
	for _, block := range bc.Blocks {
		fmt.Printf("Prev has: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("----------------------\n")
	}
}
