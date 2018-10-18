package main

import (
	"fmt"
	"sync"
	"time"

	pb "github.com/gautamrege/gochat/api"
)

type Handle struct {
	pb.Handle
	Created_at time.Time
}

// Ensure that handles are added / removed using a mutex!
type HandleSync struct {
	sync.RWMutex
	HandleMap map[string]Handle
}

var ME pb.Handle
var HANDLES HandleSync

// Insert handle if not exists
// if existst then update the data
func (hs *HandleSync) Insert(h pb.Handle) (err error) {
	hs.Lock()
	_, ok := hs.HandleMap[h.Name]
	hs.HandleMap[h.Name] = Handle{
		Handle: pb.Handle{
			Name: h.Name,
			Port: h.Port,
			Host: h.Host,
		},
		Created_at: time.Now(),
	}
	if !ok {
		fmt.Printf("\nNew User joined the chat: @%s\n> ", h.Name)
	}
	hs.Unlock()
	return nil
}

// get the user details from the map with given name
func (hs *HandleSync) Get(name string) (h pb.Handle, ok bool) {
	hs.Lock()
	tmp, ok := hs.HandleMap[name]
	if ok {
		h.Name = tmp.Name
		h.Port = tmp.Port
		h.Host = tmp.Host
	}
	hs.Unlock()

	return
}

// delete the user from map
func (hs *HandleSync) Delete(name string) {
	hs.Lock()
	delete(hs.HandleMap, name)
	hs.Unlock()
	fmt.Println("Handle Removed for ", name)
}

func (h Handle) String() string {
	return fmt.Sprintf("%s@%s:%d", h.Name, h.Host, h.Port)
}

func (hs HandleSync) String() string {
	users := "\n"
	for name, _ := range hs.HandleMap {
		users = fmt.Sprintf("%s@%s\n", users, name)
	}

	return users
}
