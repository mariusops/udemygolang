/*
Select is used to pull values off of multiple channels.
We use case to specify which channel we want to pull a value off of.
*/

package main

import "fmt"

func main() {

	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("about to exit")

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			println("from the eve channel:", v)
		case v := <-o:
			println("from the odd channel:", v)
		case v := <-q:
			println("from the quit channel:", v)
			return
		}
	}
}
