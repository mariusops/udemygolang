/*
WaitGroup

To launch a function as a Goroutine, we use the -->go<-- keyword followed by the function call.
The main function is also a Goroutine. When the main function exits, all Goroutines are killed.

We used runtime package to get the number of CPUs and number of Goroutines.
https://pkg.go.dev/runtime

We used sync package to create a WaitGroup.
https://pkg.go.dev/sync
Variable wg is a WaitGroup. Add method is used to increment the WaitGroup counter by 1.
Wait method is used to block the Goroutine until the counter becomes zero.
Done method is used to decrement the WaitGroup counter by 1.

*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
