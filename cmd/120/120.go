package main

import (
	"fmt"
)

func main() {
	charInterests := map[string][]string{
		"bond_james":       {"shaken, not stirred"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}

	for character, interests := range charInterests {
		fmt.Println(character)
		for interest, i := range interests {
			fmt.Println("\t", interest, i)
		}
	}

}
