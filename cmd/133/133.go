// Syntax of functions in Go
package main

import (
	"fmt"
)

func init() {
	fmt.Println("init() function called")
}

func main() {
	foo()
	bar("Hello")
	fmt.Println(aloha("World"))
	fmt.Println(human("John", 30))
}

func foo() {
	fmt.Println("foo() function called")
}

func bar(s string) {
	fmt.Println("foo() function called with parameter", s)
}

func aloha(s string) string {
	return fmt.Sprint("Aloha ", s)
}

func human(name string, age int) (string, int) {
	return fmt.Sprint("Hello ", name), age
}

// func (r receiver) identifier(parameters) (returns(s)) { code }
