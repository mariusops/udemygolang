package main

import "fmt"

func main() {
	am := map[string]int{
		"Gary": 42,
		"Mary": 34,
	}

	fmt.Println(am)
	for v := range am {
		fmt.Println(v)
	}

	v, ok := am["Gary"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Not found")
	}

	if w, ok := am["Gary"]; !ok {
		fmt.Println("Not found")
	} else {
		fmt.Println("Found", w)
	}
}
