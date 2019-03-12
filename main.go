package main

import (
	"bufio"
	"dposplay/account"
	"dposplay/crypto"
	"dposplay/model"
	"dposplay/other"
	"dposplay/validator"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

// Blockchain is a series of validated Blocks
var Blockchain []model.Block

// TempBlocks ...
var TempBlocks []model.Block

// CandidateBlocks handles incoming blocks for validation
var CandidateBlocks = make(chan model.Block)

// validators keeps track of open validators and balances
var validators = make(map[string]int)

// Announcements broadcasts winning validator to all nodes
var announcements = make(chan string)

var mutex = &sync.Mutex{}

// HandleConn handles the connection
func HandleConn(conn net.Conn) {
	defer conn.Close()

	go func() {
		for {
			msg := <-announcements
			io.WriteString(conn, msg)
		}
	}()
	// validator address
	var address string

	// allow user to allocate number of tokens to stake
	// the greater the number of tokens, the greater chance to forging a new block
	io.WriteString(conn, "Enter your token balance:")
	scanBalance := bufio.NewScanner(conn)
	for scanBalance.Scan() {
		balance, err := strconv.Atoi(scanBalance.Text())
		if err != nil {
			log.Printf("%v not a number: %v", scanBalance.Text(), err)
			return
		}
		t := time.Now()
		address = crypto.CalculateHash(t.String())
		validators[address] = balance
		fmt.Println(validators)
		break
	}

	io.WriteString(conn, "\nPropose data:")

	scanData := bufio.NewScanner(conn)

	go func() {
		for {
			// take in BPM from stdin and add it to blockchain after conducting necessary validation
			for scanData.Scan() {
				data := scanData.Text()
				//validation then -> delete(validators, address)

				mutex.Lock()
				oldLastIndex := Blockchain[len(Blockchain)-1]
				mutex.Unlock()

				// create newBlock for consideration to be forged
				newBlock, err := validator.GenerateBlock(oldLastIndex, data, address)
				if err != nil {
					log.Println(err)
					continue
				}
				if validator.IsBlockValid(newBlock, oldLastIndex) {
					CandidateBlocks <- newBlock
				}
				io.WriteString(conn, "\nPropose Data:")
			}
		}
	}()

	// simulate receiving broadcast
	for {
		time.Sleep(time.Minute)
		mutex.Lock()
		//output, err := json.Marshal(Blockchain)
		s, _ := other.Marshal(Blockchain)
		fmt.Println(string(s))
		mutex.Unlock()

		io.WriteString(conn, string(s)+"\n")
	}
}

// pick a winner every 15 seconds
func pickWinner() {
	time.Sleep(15 * time.Second)
	mutex.Lock()
	temp := TempBlocks
	mutex.Unlock()

	lotteryPool := []string{}
	if len(temp) > 0 {
		// from all validators who submitted a block, weight them by the number of staked tokens
		// in traditional proof of stake, validators can participate without submitting a block to be forged
	OUTER:
		for _, block := range temp {
			// if already in lottery pool, skip
			for _, node := range lotteryPool {
				if block.Validator == node {
					continue OUTER
				}
			}

			// lock list of validators to prevent data race
			mutex.Lock()
			setValidators := validators
			mutex.Unlock()

			k, ok := setValidators[block.Validator]
			if ok {
				for i := 0; i < k; i++ {
					lotteryPool = append(lotteryPool, block.Validator)
				}
			}
		}

		// randomly pick winner from lottery pool
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		lotteryWinner := lotteryPool[r.Intn(len(lotteryPool))]

		// add block of winner to blockchain and let all the other nodes know
		for _, block := range temp {
			if block.Validator == lotteryWinner {
				mutex.Lock()
				Blockchain = append(Blockchain, block)
				mutex.Unlock()
				for range validators {
					announcements <- "\n\nround winning validator: " + lotteryWinner + "\n\n"
				}
				break
			}
		}
	}

	mutex.Lock()
	TempBlocks = []model.Block{}
	mutex.Unlock()
}

func main() {

	account.AddAccountWithBalance("0x111", 100)
	account.AddAccountWithBalance("0x222", 200)
	account.ListAllAccount()

	// create genesis block
	genesisBlock := model.Block{}
	genesisBlock = model.Block{Index: 0, Timestamp: time.Now().String(), Data: "", Hash: crypto.CalculateBlockHash(genesisBlock), PrevHash: "", Validator: ""}
	fmt.Println("GENESIS BLOCK CREATED >--------->>")
	Blockchain = append(Blockchain, genesisBlock)

	fmt.Println("Starting blockchain server on port 5000")
	// start TCP and serve TCP server
	server, err := net.Listen("tcp", ":"+"5000")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	go func() {
		for candidate := range CandidateBlocks {
			mutex.Lock()
			TempBlocks = append(TempBlocks, candidate)
			mutex.Unlock()
		}
	}()

	go func() {
		for {
			pickWinner()
		}
	}()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go HandleConn(conn)
	}

}
