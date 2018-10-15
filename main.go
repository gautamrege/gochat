package main

func main() {
	// Parse flags for host, port and name
	// Create your own Global Handle ME

	var wg sync.WaitGroup
	wg.Add(3)

	// exit channel is a buffered channel for 5 exit patterns
	exit := make(chan bool, 5)

	// Listener for is-alive broadcasts from other hosts. Listening on 33333
	go registerHandle(&wg, exit)

	// Broadcast for is-alive on 33333 with own Handle.
	go isAlive(&wg, exit)

	// Cleanup Dead Handles
	go cleanupDeadHandles(&wg, exit)

	// gRPC listener
	go listen(&wg, exit)

	for {
		// Loop indefinitely and render Term
		// When we need to exit, send true 3 times on exit channel!
	}

	// exit cleanly on waitgroup
	wg.Wait()
	close(exit)
}

// Broadcast Listener
// Listens on 33333 and updates the Global Handles list
func registerHandle(wg *sync.WaitGroup, exit chan bool) {
	defer wg.Done()
	// Check if the handle is already in HANDLES. If not, add a new one!
}

// isAlive go-routine that publishes it's Handle on 33333
func isAlive(wg *sync.WaitGroup, exit chan bool) {
	defer wg.Done()
}

// cleanup Dead Handlers
func cleanupDeadHandles(wg *sync.WaitGroup, exit chan bool) {
	defer wg.Done()
	// wait for DEAD_HANDLE_INTERVAL seconds before removing them from chatrooms and handle list
}

// gRPC listener
func listen(wg *sync.WaitGroup, exit chan bool) {
	defer wg.Done()

	// initiate gRPC
}
