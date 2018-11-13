package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/gautamrege/gochat/api"
)

//Edit: Local Network broadcast address
const broadcastAddress = "192.168.1.255:33333"

// Broadcast Listener , Listens on 33333 and updates the Global Users list
func registerUser(wg *sync.WaitGroup) {
	defer wg.Done()

	var user api.Handle
	for {
		// listen to port 33333
		localAddress, _ := net.ResolveUDPAddr("udp4", MyHandle.Broadcastaddress)
		connection, err := net.ListenUDP("udp", localAddress)
		if err != nil {
			fmt.Println(err)
		}

		// read the data and add to users.
		inputBytes := make([]byte, 4096)
		length, _, _ := connection.ReadFromUDP(inputBytes)
		buffer := bytes.NewBuffer(inputBytes[:length])
		decoder := gob.NewDecoder(buffer)
		decoder.Decode(&user)
		//Ignore the user with same host
		if user.Host != MyHandle.Host {
			USERS.Insert(user)
		}

		connection.Close()
	}
}

// isAlive go-routine that publishes it's Handle on 33333
func isAlive(wg *sync.WaitGroup, exit chan bool) {
	defer wg.Done()

	// Broadcast immediately at the start
	broadcastIsAlive()

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()
	for {
		select {
		case <-exit:
			break
		case <-ticker.C:
			broadcastIsAlive()
		}
	}
}

// broadcast on 33333 every 30 seconds with MyHandle(own) Handler
func broadcastIsAlive() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	conn, err := net.Dial("udp", MyHandle.Broadcastaddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	encoder.Encode(MyHandle)
	conn.Write(buffer.Bytes())
	buffer.Reset()
	conn.Close()
}
