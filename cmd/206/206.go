/*
Atomic
Atomic operations are operations that are performed as a single unit of work without the possibility of interference from other operations.
Atomic operations are important in concurrent programming,
as they provide a way to synchronize memory access between multiple goroutines.


We use atomic package to write to an int64 variable.
We use atomic.AddInt64() to increment the counter,
and atomic.LoadInt64() to read the counter.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64 // think package atomic when you see int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
