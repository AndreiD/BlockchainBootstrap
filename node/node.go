package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// the gRPC port the server listens on
	port = ":8080"
)

type Server struct {
	Blockchain *blockchain.Blockchain
}

func (s *Server) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(in.Data)
	return &proto.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

func (s *Server) GetBlockchain(ctx context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)
	for _, b := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash: b.PrevBlockHash,
			Hash:          b.Hash,
			Data:          b.Data,
		})
	}
	return resp, nil
}

func main() {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("error listener %s\n", err)
		os.Exit(1)
	}

	srv := grpc.NewServer()

	// register & create genesis
	proto.RegisterBlockchainServer(srv, &Server{
		Blockchain: blockchain.NewBlockchain(),
	})

	fmt.Printf(">>>> Server running on port %s >>>>\n", port)

	if err := srv.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
