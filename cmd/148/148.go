//closure is when we have “enclosed” the scope of a variable in some code block.
//you can declare a variable in a function and then return a function that can access that variable

package main

import "fmt"

func main() {

	// f is a function that returns an int
	// f got a variable x from the outer scope
	// x is enclosed in the scope of f
	// f is a closure
	// f is a function that closes over the variable x
	f := incrementor()
	// here f runs the function incrementor() and returns the inner function
	fmt.Println(f())

}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
