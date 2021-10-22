package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/gautamrege/gochat/api"
)

const helpStr = `Commands
1. /users :- Get list of live users
2. @{user} message :- send message to specified user
3. /exit :- Exit the Chat
4. /all :- Send message to all the users [TODO]`

var (
	name      = flag.String("name", "", "The name you want to chat as")
	port      = flag.Int("port", 12345, "Port that your server will run on.")
	host      = flag.String("host", "", "Host IP that your server is running on.")
	stdReader = bufio.NewReader(os.Stdin)
)

var MyHandle api.Handle
var USERS = PeerHandleMapSync{
	PeerHandleMap: make(map[string]api.Handle),
}

var TERM = TermChat{}
var WS = WsChat{}

type Chatter interface {
	Input() (string, error)
	Render(string) error
	Moderate(api.ChatRequest)
}

func main() {
	// Parse flags for host, port and name
	flag.Parse()

	// TODO-WORKSHOP-STEP-1: If the name and host are empty, return an error with help message

	// TODO-WORKSHOP-STEP-2: Initialize global MyHandle of type api.Handle
	MyHandle = api.Handle{
		Name: *name,
		Host: *host,
		Port: int32(*port),
	}

	quit := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(3)

	// Broadcast for is-alive on 33333 with own UserHandle.
	go broadcastOwnHandle(&wg)

	// Listener for is-alive broadcasts from other hosts. Listening on 33333
	go listenAndRegisterUsers(&wg)

	// gRPC listener
	go startServer(&wg)

	// WebSocket listener
	go WSRun()

	// Moderate chat
	go moderateChat(quit)

	// Term ChatRoom
	for {
		textInput, err := TERM.Input()
		if err != nil {
			TERM.Render("Unable to get input.. exiting!")
			os.Exit(1) // ??
		}
		parseAndExecInput(&TERM, "term", textInput)
	}
}

// Handle the input chat messages as well as help commands
func parseAndExecInput(chat Chatter, source, input string) {
	// Split the line into 2 tokens (cmd and message)
	tokens := strings.SplitN(input, " ", 2)
	cmd := tokens[0]

	switch {
	case cmd == "":
		break
	case cmd == "?":
		chat.Render(helpStr)
		break
	case strings.ToLower(cmd) == "/users":
		chat.Render(USERS.String())
		break
	case strings.ToLower(cmd) == "/exit":
		os.Exit(1)
		break
	case cmd[0] == '@':
		// TODO-WORKSHOP-STEP-9: Write code to sendChat. Example
		// "@gautam hello golang" should send a message to handle with name "gautam" and message "hello golang"
		// Invoke sendChat to send the  message
		recvHandle, ok := USERS.Get(cmd[1:len(cmd)])
		if !ok {
			chat.Render(fmt.Sprintf("No such user: %s", cmd))
			break
		}

		sendChat(recvHandle, source, tokens[1])
		break
	case strings.ToLower(cmd) == "/help":
		chat.Render(helpStr)
		break
	default:
		chat.Render(helpStr)
	}
}
