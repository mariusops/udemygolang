/*
See if you can get a basic example of testing working
*/

package main

import (
	"fmt"
)

func main() {
	number := add(2, 3)
	fmt.Println(number)
}

func add(x, y int) int {
	return x + y
}
