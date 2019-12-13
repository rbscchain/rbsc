package main

import (
	"../rbsc/BLC"
	"fmt"
)
func main() {
	//block := BLC.NewBlock("Genenis Block", 1, []byte{0, 0, 0,0,0,0,0, 0, 0,0,0,0,0, 0, 0,0,0,0,0, 0, 0,0,0,0,0, 0, 0,0,0,0,0,0})
	//genesisBlock :=BLC.CreateBlockchainWithGenesisBlock()
	//fmt.Println(genesisBlock)
	//创世区块
	blockchain:=BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println(blockchain)
	blockchain.PrintChain()
	defer blockchain.DB.Close()
	//新增区块
	//blockchain.AddBlockToBlockchain("Send 100rbsc to xiaoming",blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	//fmt.Println(blockchain.Blocks)
	//block:=BLC.NewBlock("Send 100rbsc to xiaoming",1,[]byte{0, 0, 0,0,0,0,0, 0, 0,0,0,0,0, 0, 0,0,0,0,0, 0, 0,0,0,0,0, 0, 0,0,0,0,0,0})
	//bytes:=block.Serialize()
	//fmt.Println(bytes)
	//创建或打开数据库
	//db,err:=bolt.Open("my.db",0600,nil)
	//if err!=nil{
	//	log.Fatal(err)
	//}
	//defer db.Close()

	////创建表
	//err=db.Update(func(tx *bolt.Tx) error{
	//	//创建BlockBucket表
	//	b,err:=tx.CreateBucket([]byte("BlockBucket"))
	//	if err!=nil{
	//		return fmt.Errorf("create bucket:%s",err)
	//	}
	//	//往表里存粗数据
	//	if b!=nil{
	//		err:=b.Put([]byte("l"),[]byte(" send 100 rbsc to xiaoming"))
	//		if err!=nil{
	//			log.Panic("数据存储失败")
	//		}
	//	}
	//	return nil
	//})
	////更新失败
	//if err!=nil{
	//	log.Panic(err)
	//}

	//查看数据
	//err=db.View(func(tx *bolt.Tx) error{
	//	//创建BlockBucket表
	//	b:=tx.Bucket([]byte("BlockBucket"))
	//	//往表里存粗数据
	//	if b!=nil{
	//		data:=b.Get([]byte("l"))
	//		fmt.Printf("%s",data)
	//	}
	//	return nil
	//})
	////查看失败
	//if err!=nil{
	//	log.Panic(err)
	//}

}
