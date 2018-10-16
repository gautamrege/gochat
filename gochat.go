package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"

	pb "github.com/gautamrege/gochat/api"
)

type chatServer struct {
}

func (s *chatServer) Chat(ctx context.Context, req *pb.ChatRequest) (res *pb.ChatResponse, err error) {
	// Find the ChatRequest.to in req
	// Dial the gRPC connection to that Handle.port
	// invoke Chat method.

	return &pb.ChatResponse{}, errors.New("")
}

// gRPC listener
func listen(wg *sync.WaitGroup, exit chan bool) {
	defer wg.Done()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGoChatServer(grpcServer, &chatServer{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
