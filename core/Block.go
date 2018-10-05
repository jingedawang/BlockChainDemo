package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {

	// Block header
	Index int64
	Timestamp int64
	PrevBlockHash string
	Hash string

	// Block data
	Data string
}

func createBlock(prevBlock Block, data string) Block{
	newBlock := Block{}
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PrevBlockHash = prevBlock.Hash
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func calculateHash(block Block) string {
	toBeHashed := string(block.Index) + string(block.Timestamp) + block.PrevBlockHash + block.Data
	hashInBytes := sha256.Sum256([]byte(toBeHashed))
	return hex.EncodeToString(hashInBytes[:])
}