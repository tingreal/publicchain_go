package BLC
import(
  "time"
//  "bytes"
)
type Block struct {
	//1.区块高度int64
  Height        int64
	//2.上一个区块  hash []byte
  PrevBlockHash []byte
	//3.交易数据 字节数组 []byte
  Data          []byte
	//4.时间戳 int64
  Timestamp     int64
	//5.hash []byte
	Hash          []byte
}
//1.创建新的区块 传入Data、
//height及前段hash(有数据库不用传，为了演示)
//返回区块
func NewBlock(data string,height int64,prevBlockHash []byte) *Block {
  return &Block{Height:height,PrevBlockHash:prevBlockHash,Data:[]byte(data),Timestamp:time.Now().Unix(),Hash:nil}
//	return &Block{height,prevBlockHash,[]byte(data),time.Now().Unix(),nil}

}
