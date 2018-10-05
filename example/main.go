package main

import (
	"BlockChainDemo/core"
)

func main()  {
	blockChain := core.CreateBlockChain()
	blockChain.AddTransaction("Send 1 BTC to Faye.")
	blockChain.AddTransaction("Send 2 BTC to Liling.")
	blockChain.Print()
}