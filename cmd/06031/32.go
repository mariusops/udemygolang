package main

import "fmt"

func main() {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
	}

	// Check if the key "banana" exists in the map
	value, ok := m["banana"]
	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found")
	}
}
