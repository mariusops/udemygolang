package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	for _, v := range array {
		fmt.Println(v, fmt.Sprintf("%T", v))
	}
}
