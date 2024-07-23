/*
Writing documentation
When you have a lot to say about a package, you can create a doc.go file in the package directory.
This file should contain a package comment that introduces the package and provides a high-level overview of what the package does.

*/

// Package 243 is an example package for the Go Workshop
package main

// Import the fmt package
import (
	"fmt"
)

// main is the entry point for the application
func main() {
	number := add(2, 3)
	fmt.Println(number)
}

// add takes two integers and returns their sum
func add(x, y int) int {
	return x + y
}
