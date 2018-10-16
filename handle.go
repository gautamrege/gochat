package main

import (
	"fmt"
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

var ME Handle
var HANDLES HandleSync

func (hs *HandleSync) Insert(h Handle) (err error) {
	_, ok := hs.HandleMap[h.Name]
	if !ok {
		fmt.Println("New Handle Register for", h.Name)
	}
	hs.Lock()
	hs.HandleMap[h.Name] = h
	hs.Unlock()
	return nil
}

func (hs *HandleSync) Delete(h Handle) {
	hs.Lock()
	delete(hs.HandleMap, h.Name)
	hs.Unlock()
	fmt.Println("Handle Removed for", h.Name)
}

func (h *Handle) String() string {
	return fmt.Sprintf("%s@%s:%d", h.Name, h.Host, h.Port)
}
