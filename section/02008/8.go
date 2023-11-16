package main

import (
	"fmt"
)

// https://pkg.go.dev/fmt

func main() {
	const name = "Marius"
	const age = 31

	// %v prints the value of the variable
	fmt.Printf("Hello, world!\nMy name is %v, and I am %v years old.\n%v is a %T, and %v is an %T.\n\n", name, age, name, name, age, age)

	// %T prints the type of the variable
	// %s prints the value of the variable
	// %d prints the value of the variable as a decimal
	fmt.Printf("Hello, world!\nMy name is %s, and I am %d years old.\n%s is a %T, and %d is an %T.", name, age, name, name, age, age)
}
