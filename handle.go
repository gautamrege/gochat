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
	Handles []Handle
}

var ME Handle
var HANDLES HandleSync

func (hs *HandleSync) Insert(h Handle) (err error) {
	return nil
}

func (hs *HandleSync) Delete(h Handle) {
}

func (h *Handle) String() string {
	return fmt.Sprintf("%s@%s:%d", h.Name, h.Host, h.Port)
}
