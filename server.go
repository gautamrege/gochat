package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/gautamrege/gochat/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type chatServer struct {
	api.UnimplementedGoChatServer
}

func (s *chatServer) Chat(ctx context.Context, req *api.ChatRequest) (res *api.ChatResponse, err error) {

	message, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Unable to marshal chat request: ", err)
		return nil, err
	}

	if req.Source == "term" {
		TERM.Render(string(message))
		TERM.Moderate(*req)
	} else if req.Source == "ws" {
		WS.Render(string(message))
		WS.Moderate(*req)
	}

	// TODO-WORKSHOP-STEP-7: If this is a chat from an unknown user, insert into PeerHandleMap
	if _, ok := USERS.Get(req.From.Name); !ok {
		USERS.Insert(*req.From)
	}
	return &api.ChatResponse{}, nil
}

// gRPC listener - register and start grpc server
func startServer(wg *sync.WaitGroup) {
	defer wg.Done()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", MyHandle.Host, MyHandle.Port))
	if err != nil {
		log.Fatalf("failed to startServer: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterGoChatServer(grpcServer, &chatServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
