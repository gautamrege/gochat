package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/gautamrege/gochat/api"
)

type chatServer struct {
}

func (s *chatServer) Chat(ctx context.Context, req *pb.ChatRequest) (res *pb.ChatResponse, err error) {
	fmt.Printf("\n%s\n> ", fmt.Sprintf("@%s says: \"%s\"", req.From.Name, req.Message))

	// TODO-WORKSHOP: If this is a chat from an unknown user, insert into HANDLES

	return &pb.ChatResponse{}, nil
}

// gRPC listener
// - register and start grpc server
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

	/****
	   // THIS CODE IS FOR REFERENCE ONLY FROM THE pb PACKAGE. DO NOT UNCOMMENT
	   type pb.ChatRequest struct {
	   	From    *pb.Handle
	   	To      *pb.Handle
	   	Message string
	   }
	*****/

	var req pb.ChatRequest
	// TODO-WORKSHOP: Create req struct of type pb.ChatRequest to send to client.Chat method

	_, err = client.Chat(ctx, &req)
	if err != nil {
		log.Printf("ERROR: Chat(): %v", err)
		HANDLES.Delete(h.Name)
	}
	return
}
