package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	pb "github.com/gautamrege/gochat/api"
)

var (
	name = flag.String("name", "", "The name you want to chat as")
	port = flag.Int("port", 12345, "Port that your server will run on.")
	host = flag.String("host", "", "Host IP that your server is running on.")
)

func main() {
	// Parse flags for host, port and name
	flag.Parse()

	if *name == "" || *host == "" {
		fmt.Println("fuck off if you don't have a name and IP address :D")
		os.Exit(1)
	}
	// Create your own Global Handle ME
	var wg sync.WaitGroup
	wg.Add(4)

	HANDLES.HandleMap = make(map[string]Handle)

	// exit channel is a buffered channel for 5 exit patterns
	exit := make(chan bool, 5)

	// Listener for is-alive broadcasts from other hosts. Listening on 33333
	go registerHandle(&wg, exit)

	// Broadcast for is-alive on 33333 with own Handle.
	go isAlive(&wg, exit)

	// Cleanup Dead Handles
	go cleanupDeadHandles(&wg, exit)

	// gRPC listener
	go listen(&wg, exit)

	ME = pb.Handle{
		Name: *name,
		Host: *host,
		Port: int32(*port),
	}

	//addFakeHandles()

	var input string
	for {
		fmt.Printf("> ")
		input = readInput()

		parseAndExecInput(input)
		// Loop indefinitely and render Term
		// When we need to exit, send true 3 times on exit channel!
	}
	//time.Sleep(1 * time.Second)
	//testChat()

	// exit cleanly on waitgroup
	wg.Wait()
	close(exit)
}

func parseAndExecInput(input string) {
	helpStr := `/users :- Get list of live users
@{user} message :- send message to specified user
/exit :- Exit the Chat`
	// Split the line into 2 tokens (cmd and message)
	tokens := strings.SplitN(input, " ", 2)
	cmd := tokens[0]
	switch {
	case cmd == "":
		break
	case cmd == "?":
		fmt.Printf(`/users : List all users
/exit  : Exit chat
@<user> Type some message. e.g. @joe This works!
\n`)
		break
	case strings.ToLower(cmd) == "/users":
		fmt.Println(HANDLES)
		break
	case strings.ToLower(cmd) == "/exit":
		os.Exit(1)
		break
	case cmd[0] == '@':
		message := "hi" // default
		if len(tokens) > 1 {
			message = tokens[1]
		}

		// send message to particular user
		if h, ok := HANDLES.Get(cmd[1:]); ok {
			sendChat(h, message)
		} else {
			fmt.Println("No user: ", cmd)
		}
		break
	case strings.ToLower(cmd) == "/help":
		fmt.Println(helpStr)
	default:
		fmt.Println(helpStr)
	}
}

// Broadcast Listener
// Listens on 33333 and updates the Global Handles list
func registerHandle(wg *sync.WaitGroup, exit chan bool) {
	defer wg.Done()
	// Check if the handle is already in HANDLES. If not, add a new one!

	handle := Handle{}
	for {
		localAddress, _ := net.ResolveUDPAddr("udp", "192.168.1.255:33333")
		connection, err := net.ListenUDP("udp", localAddress)
		if err != nil {
			fmt.Println(err)
		}
		inputBytes := make([]byte, 4096)
		length, _, _ := connection.ReadFromUDP(inputBytes)
		buffer := bytes.NewBuffer(inputBytes[:length])
		decoder := gob.NewDecoder(buffer)
		decoder.Decode(&handle)
		if handle.Host != *host {
			//fmt.Println("listened data", handle)
			HANDLES.Insert(handle.Handle)
		}
		connection.Close()
	}
}

// isAlive go-routine that publishes it's Handle on 33333
func isAlive(wg *sync.WaitGroup, exit chan bool) {
	defer wg.Done()

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	for {
		select {
		case <-exit:
			break
		default:
			conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{192, 168, 1, 255}, Port: 33333})
			if err != nil {
				fmt.Println(err)
			}
			defer conn.Close()
			handle := Handle{
				Handle: pb.Handle{
					Name: *name,
					Port: int32(*port),
					Host: *host,
				},
				Created_at: time.Now(),
			}

			//fmt.Println("Broadcast: ", handle)
			encoder.Encode(handle)
			conn.Write(buffer.Bytes())
			buffer.Reset()
			time.Sleep(time.Second * 10)
		}
	}
}

// cleanup Dead Handlers
func cleanupDeadHandles(wg *sync.WaitGroup, exit chan bool) {
	defer wg.Done()
	// wait for DEAD_HANDLE_INTERVAL seconds before removing them from chatrooms and handle list
}

func testChat(message string) {
	h := pb.Handle{
		Name: "Anuj",
		Host: "192.168.1.18",
		Port: int32(3000),
	}
	sendChat(h, message)
}

func addFakeHandles() {
	for i := 0; i < 10; i++ {
		h := pb.Handle{
			Name: fmt.Sprintf("test+%d", i),
			Port: int32(i * 23),
			Host: "fake IP",
		}
		HANDLES.Insert(h)
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	return text
}
