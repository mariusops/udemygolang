// dereferencing pointers

package main

import "fmt"

func main() {

	x := 42
	fmt.Println(x)

	// & gives the address of the value
	fmt.Println(&x)

	y := &x

	*y = 43

	fmt.Println(x)

}
