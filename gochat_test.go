package main

import (
	"fmt"

	pb "github.com/gautamrege/gochat/api"
)

func addFakeHandles() {
	for i := 0; i < 10; i++ {
		h := pb.Handle{
			Name: fmt.Sprintf("test+%d", i),
			Port: int32(i * 23),
			Host: "fake IP",
		}
		HANDLES.Insert(h)
	}
}
