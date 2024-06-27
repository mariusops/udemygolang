/*
Race condition: when two or more threads access shared data and try to change it at the same time.
The result of the change depends on the order of execution of the threads.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	// GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting.
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("CPU: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	counter := 0

	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second * 2)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	fmt.Println("Counter: ", counter)
}
