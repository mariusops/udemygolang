// func expression
// type function
// first-class citizen
// assign a function to a variable
// produces a single value

package main

import "fmt"

func main() {
	fmt.Println("Hello")
	foo()

	// anonymous function
	// function without a name

	x := func() {
		fmt.Println("Anonymous function")
	}

	// anonymous function with parameters
	y := func(x int) {
		fmt.Println("The meaning of life:", x)
	}

	x()

	y(42)

}

func foo() {
	fmt.Println("foo")
}
