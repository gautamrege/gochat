package main

import (
	"fmt"

	"github.com/gautamrege/gochat/api"
)

func addFakeHandles() {
	for i := 0; i < 10; i++ {
		h := api.Handle{
			Name: fmt.Sprintf("test+%d", i),
			Port: int32(i * 23),
			Host: "fake IP",
		}
		USERS.Insert(h)
	}
}
