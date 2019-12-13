package BLC

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"

	//_ "github.com/boltdb/bolt"
)

const dbName = "rbsc.db"

const blockTableName = "blocks"

type Blockchain struct {
	//Blocks []*Block //存储有序的区块
	Tip []byte //最新的区块的hash
	DB  *bolt.DB
}

//遍历输出所有区块的信息
func (blockchain *Blockchain) PrintChain() {
	var block *Block
	var currentHash []byte = blockchain.Tip
	for {
		err := blockchain.DB.View(func(tx *bolt.Tx) error {
			//1、获取表
			b := tx.Bucket([]byte(blockTableName))
			if b != nil {
				//2、获取当前区块的字节数据
				blockBytes := b.Get(currentHash)
				block = DeserializeBlock(blockBytes)
				fmt.Printf("Height:%d\n", block.Height)
				fmt.Printf("PrevBlockHash:%x\n", block.PrevBlockHash)
				fmt.Printf("Data:%s\n", block.Data)
				fmt.Printf("Timestamp:%d\n", block.Timestamp)
				fmt.Printf("Hash:%x\n", block.Hash)
				fmt.Printf("Nonce:%d\n", block.Nonce)
			}
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
		currentHash = block.PrevBlockHash
	}
}

//增加区块到区块链里面
func (blockchain *Blockchain) AddBlockToBlockchain(data string) {
	//创建或打开数据库
	err := blockchain.DB.Update(func(tx *bolt.Tx) error {
		//1、获取表
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			//2、获取最新区块的信息
			blockBytes := b.Get(blockchain.Tip)
			//反序列化
			block := DeserializeBlock(blockBytes)
			//3、创建新的区块
			newBlock := NewBlock(data, block.Height+1, block.Hash)
			//4、将区块序列化存储到数据库当中
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			//5、更新数据库里面的"l"对应的hash
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			//6、更新blockchain的Tip
			blockchain.Tip = newBlock.Hash
		}

		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}
}

//1、创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	//创建或打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	var blockHash []byte
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			//创建BlockBucket表
			b, err = tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				return fmt.Errorf("create bucket:%s", err)
			}
		}
		//往表里存粗数据
		if b != nil {
			//创建创世区块
			genesisBlock := CreateGenesisBlock("Genesis Block")
			//err:=b.Put([]byte("l"),[]byte(" send 100 rbsc to xiaoming"))
			//存储创世区块
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic("数据存储失败")
			}
			//存储最新的区块的hash
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic("数据存储失败")
			}
			blockHash = genesisBlock.Hash

		}
		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}

	//返回区块链对象
	//return &Blockchain{[]*Block{genesisBlock}}
	return &Blockchain{blockHash, db}
}

//将区块序列化成字节数组
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

//将字节数组反序列化成区块
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

////创建或打开数据库
//db,err:=bolt.Open("my.db",0600,nil)
//if err!=nil{
//	log.Fatal(err)
//}
//defer db.Close()
//
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
//log.Panic(err)
//}
