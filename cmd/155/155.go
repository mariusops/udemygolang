// defer

package main

import "fmt"

func main() {

	// Defer order is LIFO
	defer fmt.Println("First Defer")                                 // print 4. line
	defer fmt.Println("Second Defer")                                // print 3. line
	defer fmt.Println(foo(2, 2, 3, 4, 5, 6, 7, 8, 9))                // print 2. line
	fmt.Println(bar([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})) // print 1. line
}

func foo(i ...int) int {
	xi := []int{}
	xi = append(xi, i...)

	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(i ...[]int) int {
	sum := 0
	for _, v := range i {
		for _, v2 := range v {
			sum += v2
		}
	}
	return sum
}
