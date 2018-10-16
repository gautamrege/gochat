package main

import (
	"errors"
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

	fmt.Sprintf("@%s says: \"%s\"", req.From.Name, req.Message)

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

func sendChat(h pb.Handle, message string) {

	dest := fmt.Sprintf("%s:%s", h.Host, h.Port)
	conn, err := grpc.Dial(dest, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewGoChatClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req := pb.ChatRequest{
		To: &pb.Handle{
			Name: h.Name,
			Host: h.Host,
			Port: h.Port,
		},
		From: &pb.Handle{
			Name: *name,
			Host: "TBD",
			Port: int32(*port),
		},
		Message: message,
	}

	res, err := client.Chat(ctx, &req)
	if err != nil {
		log.Fatalf("%v.Chat(_) = _, %v: ", client, err)
	}
	log.Println(res)
	return

}
