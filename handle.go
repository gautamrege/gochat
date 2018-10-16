package main

import (
	"fmt"
	pb "github.com/gautamrege/gochat/api"
	"sync"
	"time"
)

type Handle struct {
	Name       string
	Host       string
	Port       int32
	Created_at time.Time
}

// Ensure that handles are added / removed using a mutex!
type HandleSync struct {
	sync.RWMutex
	HandleMap map[string]Handle
}

var ME pb.Handle
var HANDLES HandleSync

func (hs *HandleSync) Insert(h Handle) (err error) {
	_, ok := hs.HandleMap[h.Name]
	if !ok {
		h.Created_at = time.Now()
		hs.Lock()
		hs.HandleMap[h.Name] = h
		hs.Unlock()
		fmt.Println("New Handle Register for", h.Name)
	}
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

func (hs *HandleSync) Delete(h Handle) {
	hs.Lock()
	delete(hs.HandleMap, h.Name)
	hs.Unlock()
	fmt.Println("Handle Removed for", h.Name)
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
