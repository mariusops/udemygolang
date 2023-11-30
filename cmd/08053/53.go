package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy" // Add the import statement for the puppy package
)

func main() {
	s1 := puppy.Bark()

	s2 := puppy.BigBarks()
	fmt.Println(s1)
	fmt.Println(s2)

	puppy.From12()

	// go get github.com/GoesToEleven/puppy@v1.2.0
	// go get github.com/GoesToEleven/puppy@v1.3.0
	// go get github.com/GoesToEleven/puppy@latest
}
