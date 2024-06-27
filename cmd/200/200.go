/*
Concurrency vs parallelism

If you have one core, you can't run two things at the same time. You can run them concurrently, but not in parallel.
Concurrency is executing many things at once, but only one at a time.
Parallelism is about doing lots of things at once.

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Concurrency vs parallelism")
}
