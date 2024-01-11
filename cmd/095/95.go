package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}

	fmt.Println(a)

	b := [...]string{"James", "Bond", "Shaken, not stirred"}

	fmt.Println(b)

	var c [2]int

	c[0] = 42
	c[1] = 43

	fmt.Println(c)
}
