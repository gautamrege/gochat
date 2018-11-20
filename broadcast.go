package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
	"sync"
	"time"

	"github.com/gautamrege/gochat/api"
)

//Edit: Local Network broadcast address
const broadcastAddress = "192.168.1.255:33333"

// Broadcast Listener , Listens on 33333 and updates the Global Users list
func listenAndRegisterUsers(wg *sync.WaitGroup) {
	defer wg.Done()

	var user api.Handle
	for {
		// startServer to port 33333
		udpAddress, _ := net.ResolveUDPAddr("udp4", broadcastAddress)
		udpConn, err := net.ListenUDP("udp", udpAddress)
		if err != nil {
			log.Print(err)
		}

		// read the data and add to users.
		inputBytes := make([]byte, 4096)
		length, _, _ := udpConn.ReadFromUDP(inputBytes)
		buffer := bytes.NewBuffer(inputBytes[:length])
		decoder := gob.NewDecoder(buffer)
		decoder.Decode(&user)

		// Ignore the user with same host
		if user.Host != MyHandle.Host {
			USERS.Insert(user)
		}

		udpConn.Close()
	}
}

// broadcastOwnHandle go-routine that publishes it's Handle on 33333
func broadcastOwnHandle(wg *sync.WaitGroup) {
	defer wg.Done()

	// Broadcast immediately at the start
	broadcastIsAlive()

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			broadcastIsAlive()
		}
	}
}

// broadcast on 33333 every 30 seconds with MyHandle(own) Handler
func broadcastIsAlive() {
	conn, err := net.Dial("udp", broadcastAddress)
	defer conn.Close()
	if err != nil {
		log.Print(err)
		return
	}

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(MyHandle)
	conn.Write(buffer.Bytes())
	buffer.Reset()
}
