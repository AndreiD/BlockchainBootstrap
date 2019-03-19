package validator

import (
	"dposplay/crypto"
	"dposplay/model"
)

// IsBlockValid makes sure block is valid by checking index
// and comparing the hash of the previous block
func IsBlockValid(newBlock, oldBlock model.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if crypto.CalculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
