package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gautamrege/gochat/api"
	"google.golang.org/grpc"
)

func generate_uuid() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func sendChat(receiverHandle api.Handle, source, message string) {
	receiverConnStr := fmt.Sprintf("%s:%d", receiverHandle.Host, receiverHandle.Port)

	receiverConn, err := grpc.Dial(receiverConnStr, grpc.WithInsecure())
	defer receiverConn.Close()
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return
	}

	chatClient := api.NewGoChatClient(receiverConn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	/****
	   // THIS CODE IS FOR REFERENCE ONLY FROM THE pb PACKAGE. DO NOT UNCOMMENT
	   type api.ChatRequest struct {
			From    *api.Handle
			To      *api.Handle
			Message string
	   }
	*****/

	var req api.ChatRequest
	// TODO-WORKSHOP-STEP-8: Create req struct of type api.ChatRequest to send to client.Chat method
	req.From = &MyHandle
	req.To = &receiverHandle
	req.Message = message
	req.Source = source
	req.Chatid = generate_uuid()

	_, err = chatClient.Chat(ctx, &req)
	if err != nil {
		log.Printf("ERROR: Chat(): %v", err)
		USERS.Delete(receiverHandle.Name)
	}
	return
}
