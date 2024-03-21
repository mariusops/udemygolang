// pointers
// post office boxes with addresses

package main

import "fmt"

func main() {

	x := 42
	fmt.Println(x)

	// & gives the address of the value
	fmt.Println(&x)

	// y is a pointer to an int
	y := &x

	// * gives the value at the address
	fmt.Println(*y)

}
