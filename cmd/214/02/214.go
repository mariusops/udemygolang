/*
Directional channels.

*/

package main

import "fmt"

func main() {
	// send only channel
	c := make(chan<- int, 2)

	c <- 42
	c <- 43

	// can't do this, cause it's a send only channel
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	fmt.Println("------")
	fmt.Printf("%T\n", c)

}
