/*
Hands on exercise #1
*/

package main

import (
	"fmt"
)

func main() {
	// c := make(chan int)

	// self-executing function aka func literal
	/*
		go func() {
			c <- 42
		}()
	*/

	// buffered channel
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
