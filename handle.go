package main

import (
	"fmt"
	"sync"
	"time"

	pb "github.com/gautamrege/gochat/api"
)

/**** This is the pb.Handle struct
 THIS IS FOR REFERENCE ONLY. DO NOT UNCOMMENT
 type pb.Handle struct {
	 Name string
	 Host string
	 Port int32
 }
****/
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
	// TODO-WORKSHOP: This code should insert the handle into the HandleMap
	hs.Unlock()
	return nil
}

// get the user details from the map with given name
func (hs *HandleSync) Get(name string) (h pb.Handle, ok bool) {
	hs.Lock()
	// TODO-WORKSHOP: This code should fetch the handle from the HandleMap based on the key name
	// TODO-THINK: Why is this in a Lock() method?
	hs.Unlock()

	return
}

// delete the user from map
func (hs *HandleSync) Delete(name string) {
	hs.Lock()
	// TODO-WORKSHOP: This code should remove the handle from the HandleMap based on the key name
	hs.Unlock()
	fmt.Println("Handle Removed for ", name)
}

func (h Handle) String() string {
	return fmt.Sprintf("%s@%s:%d", h.Name, h.Host, h.Port)
}

func (hs HandleSync) String() (users string) {
	// TODO-WORKSHOP: This code should print the list of all names of the handles in the map
	// TODO-THINK: Do we need a Lock here?

	return
}
