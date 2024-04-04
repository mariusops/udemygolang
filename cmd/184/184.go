// Type constraints
// Type set interface

package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

// ~ any alias with the same underlying type
type myNumber interface {
	~int | ~float64
}

func addT[T myNumber](a, b T) T {
	return a + b
}

// type alias with underlying type int
type myAlias int

func main() {
	// variable n is of type myAlias
	var n myAlias = 42

	fmt.Println(addI(1, 2))

	fmt.Println(addF(1.1, 2.2))

	fmt.Println(addT(n, 2))
	fmt.Println(addT(1.1, 2.2))
}
