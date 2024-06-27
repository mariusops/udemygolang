/*
Do not communicate by sharing memory; instead, share memory by communicating.
Only one goroutine has access to the channel at a time, so the data is always synchronized.
Gorotuines shares the same address space, so they can access the same memory.
*/

package main

import (
	"fmt"
)

func main() {
	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)
}
