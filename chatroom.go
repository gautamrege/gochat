package main

import (
	"sync"

	"github.com/gautamrege/gochat/api"
)

/* This is a chatroom that is registered on this server! */
type Chatroom struct {
	Name       string
	People     []api.Handle
	ChatStream chan Chat
}

/* This is the message that is broadcast on the wire */
type Chat struct {
	From    api.Handle
	To      api.Handle
	Message string
}

// Ensure that chatrooms are added / removed using a mutex!
type ChatroomSync struct {
	sync.RWMutex
	Chatrooms []Chatroom
}

var CHATROOMS ChatroomSync

func (hs *ChatroomSync) Insert(h Chatroom) (err error) {
	return nil
}

func (hs *ChatroomSync) Delete(h Chatroom) {
}

func NewChatroom(name string) (c *Chatroom) {
	// create a new chatroom and return it.
	return nil
}

func (c *Chatroom) Print(chat Chat) {
	// Prints chat on the terminal.
	// Based on where the chat window is opened, the term may have a positional row/col setting
}

func (c *Chatroom) Message(text string) {
	// sends a message via gRPC to all the UserHandle in that chat-room
}
