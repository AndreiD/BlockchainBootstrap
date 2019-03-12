package model

// Block is the base structure in the blockchain
type Block struct {
	Index     int    // index of the block
	Timestamp string //  the time of creation of block in milliseconds
	Data      string // data in the block or the transactions
	Hash      string // hash of the current block
	PrevHash  string // hash of the last block on the chain
	Validator string //  the address of the account whose made this block
	Signature string //  the encrypted hash of the block, signed by the validator
}
