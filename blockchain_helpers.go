package main

import (
	"strconv"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// calculates the sha256 hash of a block
func calculateHash(block Block) string {
	record := string(block.Index) + strconv.FormatInt(block.Timestamp, 10) + block.TheData + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// generates a new block
func generateBlock(oldBlock Block, TheData string, address string) (Block, error) {
	var newBlock Block
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().UnixNano()
	newBlock.TheData = TheData
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	newBlock.Validator = address
	return newBlock, nil
}

// checks if the new block is valid
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// choose the longest length block
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
