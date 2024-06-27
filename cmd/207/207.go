/*
Hands-on exercise #1
● in addition to the main goroutine, launch two additional goroutines
○ each additional goroutine should print something out
● use waitgroups to make sure each goroutine finishes before your program exists
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Println("foo")
	wg.Done()
}

func bar() {
	fmt.Println("bar")
	wg.Done()
}

func main() {

	fmt.Println("Start: Number of CPUs: ", runtime.NumCPU())
	fmt.Println("Start: Number of Goroutines: ", runtime.NumGoroutine())
	wg.Add(2)
	go foo()
	go bar()

	fmt.Println("Before wait: Number of CPUs: ", runtime.NumCPU())
	fmt.Println("Before wait: Number of Goroutines: ", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("After: Number of CPUs: ", runtime.NumCPU())
	fmt.Println("After: Number of Goroutines: ", runtime.NumGoroutine())

	fmt.Println("Both goroutines finished. Exiting...")
}
