package main

import "fmt"

func main() {

	// var for zero value
	var x int
	fmt.Println(x)

	// short declaration operator
	y := 5
	fmt.Println(y)

	// multiple short declaration operator
	a, b := 1, 2
	fmt.Println(a, b)

	// var when specificity is required
	var c float32 = 3.12
	fmt.Println(c)

	// blank identifier
	d, e, f := 41, 42, 43
	fmt.Println(d, e, f)

}
