package main

import (
	"fmt"
	pb "github.com/gautamrege/gochat/api"
	"sync"
	"time"
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

func (hs *HandleSync) Insert(h pb.Handle) (err error) {
	hs.Lock()
	_, ok := hs.HandleMap[h.Name]
	if !ok {
		hs.HandleMap[h.Name] = Handle{
			Handle: pb.Handle{
				Name: h.Name,
				Port: h.Port,
				Host: h.Host,
			},
			Created_at: time.Now(),
		}
		fmt.Println("New Handle Register for", h.Name)
	}
	hs.Unlock()
	return nil
}

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

func (hs *HandleSync) Delete(name string) {
	hs.Lock()
	delete(hs.HandleMap, name)
	hs.Unlock()
	fmt.Println("Handle Removed for", name)
}

func (h Handle) String() string {
	return fmt.Sprintf("%s@%s:%d", h.Name, h.Host, h.Port)
}

func (hs HandleSync) String() string {
	users := "\n"
	for name, _ := range hs.HandleMap {
		users = fmt.Sprintf("%s%s\n", users, name)
	}

	return users
}
