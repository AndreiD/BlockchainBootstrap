package main

import (
	"flag"
	"log"
	"time"

	"github.com/AndreiD/grpsplay/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listFlag := flag.Bool("list", false, "get the blockchain")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting to server %s", err)
	}

	client = proto.NewBlockchainClient(conn)

	if *addFlag {
		addBlock()
	}
	if *listFlag {
		getBlockchain()
	}

}

func addBlock() {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if err != nil {
		log.Fatalf("could not add block %s", err)
	}

	log.Printf("new block added successfully: %s\n", block.Hash)
}

func getBlockchain() {
	block, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatalf("unable to get blockchain: %s", err)
	}

	log.Println("~~~~~~~~~ Blockchain ~~~~~~~~~")
	for _, b := range block.Blocks {
		log.Println("--------------")
		log.Printf("index: %d | timestamp: %s | data: %s | hash: %s | prev hash: %s\n", b.Index, b.Timestamp, b.Data, b.Hash, b.PrevBlockHash)
		log.Println("--------------")
	}
	log.Println("~~~~~~~~~ End Blockchain ~~~~~~~~~")
}
