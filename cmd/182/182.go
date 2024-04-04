// Generics
// Type constraints
// DRY code - Don't Repeat Yourself
// https://golang.org/doc/effective_go.html#generics
package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

// The type constraint is defined in the function signature
// T is constrained to int or float64
func addT[T int | float64](a, b T) T {
	return a + b
}

func main() {

	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.1, 2.2))
	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.1, 2.2))
}
