package main

import (
	"fmt"
	"time"
)

// A function to simulate some work
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// 1. Goroutines
	// The 'go' keyword starts a new goroutine (lightweight thread).

	// This runs in the background
	go say("world")

	// This runs in the main thread
	say("hello")

	// Note: If main() finishes, all goroutines are killed immediately.
	// We usually use WaitGroups or Channels to wait for them.

	// 2. Channels
	// Channels are typed conduits for sending and receiving values.

	messages := make(chan string)

	// Start a goroutine that sends a message to the channel
	go func() {
		messages <- "ping" // Send "ping" to channel
	}()

	// Receive the message from the channel
	msg := <-messages
	fmt.Println("Received:", msg)

	// 3. Buffered Channels
	// Can hold a limited number of values without a receiver ready.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Buffered:", <-ch)
	fmt.Println("Buffered:", <-ch)
}
