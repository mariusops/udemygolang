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

// The type is defined in interface
type myNumber interface {
	int | float64
}

// Here the type constraint from the interface is used
func addT[T myNumber](a, b T) T {
	return a + b
}

func main() {

	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.1, 2.2))
	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.1, 2.2))
}
