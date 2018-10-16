package blockchain

import (
	"crypto/sha1"
	"encoding/hex"
	"time"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// Block is the base block...
type Block struct {
	Index         int64
	TimeStamp     *timestamp.Timestamp
	Hash          string
	PrevBlockHash string
	Data          string
}

// Blockchain is a list of blocks
type Blockchain struct {
	Blocks []*Block
}

// calculates the hash. sha1 is used because it produces shorter hashes
// for our educational purposes it suffice
func (b *Block) setHashAndTimestamp() {
	hash := sha1.Sum([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
	b.TimeStamp = &timestamp.Timestamp{Seconds: time.Now().Unix()}
}

// Create a new block
func NewBlock(index int64, data string, prevBlockHash string) *Block {
	block := &Block{
		Index:         index,
		Data:          data,
		PrevBlockHash: prevBlockHash,
		Hash:          "",
	}
	block.setHashAndTimestamp()
	return block
}

// AddBlock adds a new block
func (bc *Blockchain) AddBlock(data string) *Block {

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, data, prevBlock.Hash)

	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}

// NewBlockchain creates a new blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewBlock(0, "~~~ Genesis Block Created ~~~", "")}}
}
