// func, named foo, returns an int
// func, named bar, returns an int and a string
// call both
// print the results

package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "James Bond"
}
