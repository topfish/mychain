package metadata

import(
	"math"
)

const (
	Dif		= 20
	INT64_MAX 	= math.MaxInt64
)

type Block struct{
	PrevHash	[]byte
	Hash		[]byte
	Data		string
	Height		int64
	Timestamp	int64
	Nonce		int
}

type BlockChain struct{
	Blocks []Block
}


