/*
Just comma ok idiom
*/
package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		c <- 42

	}()

	v, ok := <-c

	fmt.Println(v, ok)

	fmt.Println("about to exit")

}
