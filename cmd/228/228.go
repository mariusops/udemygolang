/*
Hands on exercise #7
*/
package main

import (
	"fmt"
)

func main() {
	goroutines := 10
	numberOfNumbers := 10
	c := make(chan int)

	for i := 0; i < goroutines; i++ {
		go func() {
			for j := 0; j < numberOfNumbers; j++ {
				c <- j
			}
		}()
	}

	for k := 0; k < goroutines*numberOfNumbers; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")

}
