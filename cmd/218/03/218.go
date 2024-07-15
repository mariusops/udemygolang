/*
Just comma ok idiom
*/
package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c

	fmt.Println(v, ok)
	fmt.Println(<-c)

	v, ok = <-c

	fmt.Println(v, ok)
	fmt.Println(<-c)

	fmt.Println("about to exit")

}
