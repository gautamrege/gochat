package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	name = flag.String("name", "", "The name you want to chat as")
	port = flag.Int("port", 12345, "Port that your server will run on.")
	host = flag.String("host", "", "Host IP that your server is running on.")
)

func main() {
	// Parse flags for host, port and name
	flag.Parse()

	// TODO-WORKSHOP: If the name and host are empty, return an error with help message

	// TODO-WORKSHOP: Initialize the HANDLES.HandleMap below using make
	// HANDLES.HandleMap = ??????

	// exit channel is a buffered channel for 5 exit patterns
	exit := make(chan bool, 5)

	var wg sync.WaitGroup

	// Listener for is-alive broadcasts from other hosts. Listening on 33333
	go registerHandle(&wg, exit)

	// Broadcast for is-alive on 33333 with own Handle.
	go isAlive(&wg, exit)

	// Cleanup Dead Handles
	go cleanupDeadHandles(&wg, exit)

	// gRPC listener
	go listen(&wg, exit)

	// TODO-WORKSHOP: Initialize global ME of type pb.Handle

	var input string
	for {
		// Accept chat input
		fmt.Printf("> ")
		input = readInput()

		parseAndExecInput(input)
	}

	// exit cleanly on waitgroup
	close(exit)
}

// Handle the input chat messages as well as help commands
func parseAndExecInput(input string) {
	helpStr := `/users :- Get list of live users
@{user} message :- send message to specified user
/exit :- Exit the Chat
/all :- Send message to all the users [TODO]`

	// Split the line into 2 tokens (cmd and message)
	tokens := strings.SplitN(input, " ", 2)
	cmd := tokens[0]
	switch {
	case strings.ToLower(cmd) == "/users":
		fmt.Println(HANDLES)
		break
	case strings.ToLower(cmd) == "/exit":
		os.Exit(1)
		break
	case cmd[0] == '@':
		// TODO-WORKSHOP: Write code to sendChat. Example
		// "@gautam hello golang" should send a message to handle with name "gautam" and message "hello golang"
		// Invoke sendChat to send the  message
		break
	case strings.ToLower(cmd) == "/help":
		fmt.Println(helpStr)
		break
	default:
		fmt.Println(helpStr)
	}
}

// cleanup Dead Handlers
func cleanupDeadHandles(wg *sync.WaitGroup, exit chan bool) {
	defer wg.Done()
	// wait for DEAD_HANDLE_INTERVAL seconds before removing them from chatrooms and handle list
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	return text
}
