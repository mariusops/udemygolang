// variadic parameters
package main

import (
	"fmt"
)

func main() {
	n1 := sum(3, 4, 5)
	fmt.Println(n1)

}

// ...type becomes a slice of that type
func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)
	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum
}

// func (r receiver) identifier(parameters) (returns(s)) { code }
