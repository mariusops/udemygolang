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

	// print the type of y
	fmt.Printf("%T\n", y)

	// z is the value of the y pointer, which is the address of x
	z := *y

	fmt.Println(z)

}
