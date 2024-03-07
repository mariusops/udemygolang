// defer

package main

import "fmt"

func main() {
	defer fmt.Println(foo(2, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(bar([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}))
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
