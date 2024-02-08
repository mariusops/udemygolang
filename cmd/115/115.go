package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":  42,
		"Mac":   24,
		"Henry": 32,
	}

	fmt.Println(am)
	fmt.Println("The age of Todd was", am["Todd"])
}
