// recursion
// A function that calls itself

package main

import "fmt"

func main() {
	// 4!
	fmt.Println(4 * 3 * 2 * 1)
	x := factorial(4)
	fmt.Println(x)

	y := factLoop(4)
	fmt.Println(y)
}

func factorial(n int) int {
	fmt.Println("n is", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// whatever you can do with recursion you can do with a loop
func factLoop(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
