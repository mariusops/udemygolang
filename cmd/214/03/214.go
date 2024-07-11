/*
Directional channels.

*/

package main

import "fmt"

func main() {
	// receive only channel
	c := make(<-chan int, 2)

	// can't do this, cause it's a receive only channel
	// c <- 42
	// c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("------")
	fmt.Printf("%T\n", c)

}
