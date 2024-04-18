// Sort
package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 33, 124, 25, 3, 9, 1, 2, 6, 8, 5, 10}
	xs := []string{"James", "Q", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)

}
