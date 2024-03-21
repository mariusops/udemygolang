// pointers, values, the stack & the heap

// the stack: fast, lightweight, used for local variables
// the stack is self cleaning
// more efficient, less bugs

// the heap: slower, more memory, used for reference types
// the heap is garbage collected
// less efficient, more bugs

package main

import "fmt"

func main() {

	// a is stored on the stack
	a := 42
	fmt.Println(a)

	// b is a pointer to an int
	// b is stored on the stack
	b := &a

	// *b is the value at the address
	// *b is stored on the heap
	fmt.Println(*b)

}
