/*

Context is a package that provides functions to manage multiple goroutines.
It provides a way to pass a context to a goroutine and cancel the context when needed.

Using cancel function, we can cancel the context and all the goroutines that are using this context will be notified.

*/

package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Context: ", ctx)
	fmt.Println("Context Err: ", ctx.Err())
	fmt.Printf("Context type:\t%T\n", ctx)
	fmt.Println("-----------------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("Context: ", ctx)
	fmt.Println("Context Err: ", ctx.Err())
	fmt.Printf("Context type:\t%T\n", ctx)
	fmt.Println("Cancel: ", cancel)
	fmt.Printf("Cancel type:\t%T\n", cancel)
	fmt.Println("-----------------")

	cancel()

	fmt.Println("Context: ", ctx)
	fmt.Println("Context Err: ", ctx.Err())
	fmt.Printf("Context type:\t%T\n", ctx)
	fmt.Println("Cancel: ", cancel)
	fmt.Printf("Cancel type:\t%T\n", cancel)
	fmt.Println("-----------------")

}
