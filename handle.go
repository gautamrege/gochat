package main

import (
	"fmt"
	"sync"
	"time"

	pb "github.com/gautamrege/gochat/api"
)

type User struct {
	pb.Handle
	Created_at time.Time
}

// Ensure that users are added / removed using a mutex!
type UserSync struct {
	sync.RWMutex
	UserMap map[string]User
}

var ME pb.Handle
var USERS UserSync

// Insert user if not exists
// if existst then update the data
func (hs *UserSync) Insert(h pb.Handle) (err error) {
	hs.Lock()
	_, ok := hs.UserMap[h.Name]
	hs.UserMap[h.Name] = User{
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
func (hs *UserSync) Get(name string) (h pb.Handle, ok bool) {
	hs.Lock()
	tmp, ok := hs.UserMap[name]
	if ok {
		h.Name = tmp.Name
		h.Port = tmp.Port
		h.Host = tmp.Host
	}
	hs.Unlock()

	return
}

// delete the user from map
func (hs *UserSync) Delete(name string) {
	hs.Lock()
	delete(hs.UserMap, name)
	hs.Unlock()
	fmt.Println("User Removed for ", name)
}

func (h User) String() string {
	return fmt.Sprintf("%s@%s:%d", h.Name, h.Host, h.Port)
}

func (hs UserSync) String() string {
	users := "\n"
	for name, _ := range hs.UserMap {
		users = fmt.Sprintf("%s@%s\n", users, name)
	}

	return users
}
