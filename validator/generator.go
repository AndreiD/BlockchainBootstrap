package validator

import (
	"dposplay/crypto"
	"dposplay/model"
	"time"
)

// GenerateBlock creates a new block using previous block's hash
func GenerateBlock(oldBlock model.Block, data string, address string) (model.Block, error) {

	var newBlock model.Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = crypto.CalculateBlockHash(newBlock)
	newBlock.Validator = address

	return newBlock, nil
}
