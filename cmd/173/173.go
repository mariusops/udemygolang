// pointer & value semantics

// pointer semantics
// when you pass a pointer, you are passing a reference
// you are passing the address

// value semantics
// when you pass a value, you are making a copy
// you are passing the value

package main

import "fmt"

// pass the value, make a copy
func foo(x int) {
	x = 43
}

// pass the address, pass the reference/pointer
func bar(x *int) {
	// dereference the address to set the new value
	*x = 43
}

func main() {

	a := 42
	fmt.Println(a)
	foo(a)
	fmt.Println(a)

	bar(&a)
	fmt.Println(a)

}
