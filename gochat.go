package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/gautamrege/gochat/api"
)

type chatServer struct {
}

func (s *chatServer) Chat(ctx context.Context, req *api.ChatRequest) (res *api.ChatResponse, err error) {
	fmt.Printf("\n%s\n> ", fmt.Sprintf("@%s says: \"%s\"", req.From.Name, req.Message))

	if _, ok := USERS.Get(req.From.Name); !ok {
		USERS.Insert(*(req.From))
	}
	return &api.ChatResponse{}, nil
}

// gRPC listener - register and start grpc server
func listen(wg *sync.WaitGroup) {
	defer wg.Done()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", MyHandle.Host, MyHandle.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterGoChatServer(grpcServer, &chatServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func sendChat(toHandle api.Handle, message string) {
	destStr := fmt.Sprintf("%s:%d", toHandle.Host, toHandle.Port)

	conn, err := grpc.Dial(destStr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return
	}

	client := api.NewGoChatClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := api.ChatRequest{
		To: &toHandle,
		From:    &MyHandle,
		Message: message,
	}

	_, err = client.Chat(ctx, &req)
	if err != nil {
		log.Printf("ERROR: Chat(): %v", err)
		USERS.Delete(toHandle.Name)
	}
	return
}
