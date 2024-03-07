// callback
// passing a func as an argument
package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)
	x := doMath(1, 2, add)
	fmt.Println(x)

	y := doMath(43, 11, subtract)
	fmt.Println(y)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
