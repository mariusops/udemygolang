/*
Ranging over a channel
Closing a channel

Important to close a channel when you're done with it.
The range c will continue to pull values off the channel until it's closed.
*/

package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
