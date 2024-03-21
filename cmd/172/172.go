// pass by value
// pointers types
// reference types: slices, maps, channels, pointers, functions, interfaces
// and mutability, slice, map and pointer are mutable

package main

import "fmt"

// pass by value
func intDelta(n *int) {
	*n = 43
}

func sliceDelta(ii []int) {
	ii[0] = 99
}

func mapDelta(mm map[string]int) {
	mm["Todd"] = 33
}

func main() {

	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["Todd"] = 44
	fmt.Println(m)

	mapDelta(m)
	fmt.Println(m)

}
