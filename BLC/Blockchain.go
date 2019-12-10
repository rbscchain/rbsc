package BLC

type Blockchain struct {
	Blocks []*Block //存储有序的区块
}

//增加区块到区块链里面
func (blockchain *Blockchain) AddBlockToBlockchain(data string,height int64,prevBlockHash []byte)  {
	//创建新区块
	newBlock:=NewBlock(data,height,prevBlockHash)
	//往链里面添加新的区块
	blockchain.Blocks=append(blockchain.Blocks,newBlock)
}

//1、创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain{
	//创建创世区块
	genesisBlock:=CreateGenesisBlock("Genesis Data ......")
	//返回区块链对象
	return &Blockchain{[]*Block{genesisBlock}}
}