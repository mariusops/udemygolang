/*
Comma ok idiom with select.

We can use the comma ok idiom to check if a channel is closed.
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

func send(even, odd, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-even:
			println("the value received from the even channel", v)
		case v := <-odd:
			println("the value received from the odd channel", v)
		case i, ok := <-quit:
			if !ok {
				println("from comma ok:", i, ok)
				return
			} else {
				println("from comma ok:", i, ok)
			}
		}
	}
}
