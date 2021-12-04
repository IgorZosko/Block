package main

import (
	"fmt"
)

func main() {

	chain := InitBlockChain()

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("third Block")

	for _, block := range chain.blocks {
		fmt.Printf("Previos Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
