package main

import (
	"fmt"
)

func main() {
	x := foo()
	fmt.Println(x)

	y := bar()
	fmt.Printf("%T\n", y)
	i := y()
	fmt.Println(i)
}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 42
	}
}
