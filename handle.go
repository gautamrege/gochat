package main

import (
	"fmt"
	"sync"

	"github.com/gautamrege/gochat/api"
)

// Ensure that users are added / removed using a mutex!
type PeerHandleMapSync struct {
	sync.RWMutex
	PeerHandleMap map[string]api.Handle
}

// Insert user if not exists already then add it
func (hs *PeerHandleMapSync) Insert(newHandle api.Handle) (err error) {
	hs.Lock()
	_, ok := hs.PeerHandleMap[newHandle.Name]
	if !ok {
		hs.PeerHandleMap[newHandle.Name] = newHandle
		fmt.Printf("\nNew UserHandle joined the chat: @%s\n> ", newHandle.Name)
	}
	hs.Unlock()
	return nil
}

// get the user details from the map with given name
func (hs *PeerHandleMapSync) Get(name string) (handle api.Handle, ok bool) {
	hs.Lock()
	handle, ok = hs.PeerHandleMap[name]
	hs.Unlock()
	return
}

// delete the user from map
func (hs *PeerHandleMapSync) Delete(name string) {
	hs.Lock()
	delete(hs.PeerHandleMap, name)
	hs.Unlock()
	fmt.Println("UserHandle Removed for ", name)
}

func String(h api.Handle) string {
	return fmt.Sprintf("%s@%s:%d", h.Name, h.Host, h.Port)
}

func (hs PeerHandleMapSync) String() string {
	users := "\n"
	for name, _ := range hs.PeerHandleMap {
		users = fmt.Sprintf("%s@%s\n", users, name)
	}
	return users
}
