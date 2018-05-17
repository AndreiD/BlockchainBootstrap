package main

//content of each block
type Block struct {
	Index     int `json:"index"`
	Timestamp int64 `json:"timestamp"`
	TheData   string  `json:"thedata"` // <- holding our "data"
	Hash      string `json:"hash"`
	PrevHash  string `json:"prevhash"`
	Validator  string `json:"validator"`
}

//Data that gets added into the block
type Message struct {
	TheData string `json:"thedata"`
	NodeName string `json:"nodename"`
}
