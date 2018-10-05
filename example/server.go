package main

import (
	"BlockChainDemo/core"
	"encoding/json"
	"io"
	"net/http"
)

var blockChain core.BlockChain

func run()  {
	http.HandleFunc("/blockchain/get", blockchainGetHandler)
	http.HandleFunc("/blockchain/write", blockchainWriteHandler)
	http.ListenAndServe("localhost:8888", nil)
}

func blockchainGetHandler(writer http.ResponseWriter, request *http.Request) {
	bytes, error := json.MarshalIndent(blockChain, "", "\t")
	if error != nil {
		http.Error(writer, error.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(writer, string(bytes))
}

func blockchainWriteHandler(writer http.ResponseWriter, request *http.Request)  {
	data := request.URL.Query().Get("data")
	blockChain.AddTransaction(data)
	blockchainGetHandler(writer, request)
}

func main()  {
	blockChain = core.CreateBlockChain()
	run()
}