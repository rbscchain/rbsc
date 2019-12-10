package main

import (
	"../rbsc/BLC"
	"fmt"
)
func main() {
	//block := BLC.NewBlock("Genenis Block", 1, []byte{0, 0, 0,0,0,0,0, 0, 0,0,0,0,0, 0, 0,0,0,0,0, 0, 0,0,0,0,0, 0, 0,0,0,0,0,0})
	//genesisBlock :=BLC.CreateGenesisBlock("Genenis Block")
	//fmt.Println(genesisBlock)
	//创世区块
	blockchain:=BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println(blockchain.Blocks)
	//新增区块
	blockchain.AddBlockToBlockchain("Send 100rbsc to xiaoming",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	fmt.Println(blockchain.Blocks)

}
