package BLC

import (
	"time"
)

type Block struct {
	//1、区块高度
	Height int64
	//2、上一个区块hash
	PrevBlockHash []byte
	//3、交易数据
	Data []byte
	//4、时间戳
	Timestamp int64
	//5、Hash
	Hash []byte
	//6、Nonce
	Nonce int64
}

//将区块序列化成字节数组
//func  (block *Block) Serialize() []byte  {
//	var result bytes.Buffer
//	encoder:=gob.NewEncoder(&result)
//	err:=encoder.Encode(block)
//	if err!=nil{
//		log.Panic(err)
//	}
//	return result.Bytes()
//}

//将字节数组反序列化成区块
//func DeserializeBlock(blockBytes []byte) *Block  {
//	var block Block
//	decoder:=gob.NewDecoder(bytes.NewReader(blockBytes))
//	err:=decoder.Decode(&block)
//	if err!=nil{
//		log.Panic(err)
//	}
//	return &block
//}

//func (block *Block) SetHash() {
//	//1、Height转化为字节数组
//	heightBytes := IntToHex(block.Height)
//	//2、Timestamp转化为字节数组
//	timeString := strconv.FormatInt(block.Timestamp, 2) //2代表转化为2进制
//	timeBytes := []byte(timeString)
//	//3、拼接所有的属性  二维字节数组
//	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})
//	//4、生成HASH
//	hash := sha256.Sum256(blockBytes)
//	block.Hash = hash[:] //切片赋值
//}

//1、创建新的区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{
		Height:        height,
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
		Nonce:         0,
	}
	//设置HASH
	//block.SetHash()
	//调用工作量证明的方法并且返回有效的HASH和Nonce
	pow:=NewProofOfWork(block)
	hash,nonce:=pow.Run()
	block.Hash=hash[:]
	block.Nonce=nonce
	return block
}

//2、创建创世区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
