package core

import "fmt"

type BlockChain struct {
	Blocks []*Block
}

func CreateBlockChain() BlockChain {
	genesisBlock := createBlock(Block{Index:-1}, "I am genesis block.")
	blockChain := BlockChain{}
	blockChain.Blocks = append(blockChain.Blocks, &genesisBlock)
	return blockChain
}

func (blockChain *BlockChain) AddTransaction(data string)  {
	block := createBlock(*blockChain.Blocks[len(blockChain.Blocks) - 1], data)
	blockChain.Blocks = append(blockChain.Blocks, &block)
}

func (blockChain *BlockChain) Print()  {
	for _, block := range blockChain.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("PreBlockHash: %s\n", block.PrevBlockHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Println()
	}
}