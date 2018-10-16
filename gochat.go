package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/gautamrege/gochat/api"
)

type chatServer struct {
}

func (s *chatServer) Chat(ctx context.Context, req *pb.ChatRequest) (res *pb.ChatResponse, err error) {
	fmt.Printf("\n%s\n>", fmt.Sprintf("@%s says: \"%s\"", req.From.Name, req.Message))

	if _, ok := HANDLES.Get(req.From.Name); !ok {
		// insert new user into HANDLES
		HANDLES.Insert(*(req.From))
	}

	return &pb.ChatResponse{}, nil
}

// gRPC listener
func listen(wg *sync.WaitGroup, exit chan bool) {
	defer wg.Done()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
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

func sendChat(h pb.Handle, message string) {
	dest := fmt.Sprintf("%s:%d", h.Host, h.Port)

	conn, err := grpc.Dial(dest, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGoChatClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := pb.ChatRequest{
		To: &pb.Handle{
			Name: h.Name,
			Host: h.Host,
			Port: h.Port,
		},
		From:    &ME,
		Message: message,
	}

	_, err = client.Chat(ctx, &req)
	if err != nil {
		log.Printf("ERROR: Chat(): %v", err)
	}
	return

}
