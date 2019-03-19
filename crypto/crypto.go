package crypto

import (
	"crypto/sha256"
	"dposplay/model"
	"encoding/hex"
)

// CalculateHash - a simple SHA256 hashing function
func CalculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//CalculateBlockHash returns the hash of all block information
func CalculateBlockHash(block model.Block) string {
	record := string(block.Index) + block.Timestamp + string(block.Data) + block.PrevHash
	return CalculateHash(record)
}
