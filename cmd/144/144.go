// anonymous function
package main

import "fmt"

func main() {
	fmt.Println("Hello")
	foo()

	// anonymous function
	// function without a name

	func() {
		fmt.Println("Anonymous function")
	}()

	// anonymous function with parameters
	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42)

}

func foo() {
	fmt.Println("foo")
}
