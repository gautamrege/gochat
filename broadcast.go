package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/gautamrege/gochat/api"
)

//Edit: Local Network broadcast address
//const broadcastAddress = "192.168.1.255:33333"
const broadcastAddress = "192.168.3.255:3333"

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

// broadcast on 33333 every 10 seconds with MyHandle(own) Handler
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

func preloadUsers() error {
	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Print(err)
		return err
	}

	log.Print("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Print(err)
		return err
	}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	var users []api.Handle

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		log.Print(err)
		return err
	}

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for _, u := range users {
		USERS.Insert(u)
		log.Printf("User %s added", u.Name)
	}

	return nil
}
