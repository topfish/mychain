package utils

import (
	"bytes"
	"encoding/binary"
	"math/big"
	"mychain/metadata"
	"crypto/sha256"
	"time"
	"fmt"
)

//int转16进制切片
func IntToHex(num int64) []byte{
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	Myerr(err)
	return buff.Bytes()
}

func Myerr(err error) {
	if err != nil {
		panic(err)
	}
}

//工作量证明
func ProofOfWork(b metadata.Block, dif int) ([]byte, int){
	target := big.NewInt(1)
	target.Lsh(target, uint(256-dif))
	nonce := 0
	for ; nonce < metadata.INT64_MAX; nonce++ {
		check := bytes.Join(
		[][]byte{b.PrevHash,
	  	  []byte(b.Data),
	  	  IntToHex(b.Height),
	  	  IntToHex(b.Timestamp),
	  	  IntToHex(int64(nonce))},
		[]byte{})
		hash := sha256.Sum256(check)
		var hashInt big.Int
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(target) == -1 {
			return hash[:], nonce
		}
	}
	return []byte(""), nonce
}

//生成创世区块
func GenesisBlock(data string) metadata.BlockChain {
	var bc metadata.BlockChain
	bc.Blocks = make([]metadata.Block, 1)
	bc.Blocks[0] = metadata.Block{
		PrevHash: 	[]byte(""),
		Data:		data,
		Height:		1,
		Timestamp:	time.Now().Unix(),
	}
	bc.Blocks[0].Hash, bc.Blocks[0].Nonce = ProofOfWork(bc.Blocks[0], metadata.Dif)
	return bc
}

//挖矿（生成新区块）
func GenerateBlock(bc *metadata.BlockChain, data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	block := metadata.Block{
		PrevHash:	prevBlock.Hash,
		Data:		data,
		Height:		prevBlock.Height+1,
		Timestamp:	time.Now().Unix(),
	}
	block.Hash, block.Nonce = ProofOfWork(block, metadata.Dif)
	bc.Blocks = append(bc.Blocks, block)
}

//展示区块信息
func Print(bc metadata.BlockChain) {
	for _, i := range bc.Blocks {
		fmt.Printf("PrevHash: %x\n", i.PrevHash)
    		fmt.Printf("Hash: %x\n", i.Hash)
    		fmt.Println("Block's Data: ", i.Data)
    		fmt.Println("Current Height: ", i.Height)
    		fmt.Println("Timestamp: ", i.Timestamp)
    		fmt.Println("Nonce: ", i.Nonce)
	}
}
