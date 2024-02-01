package main

import (
	"fmt"
)

func main() {
	charInterests := make(map[string][]string)

	charInterests["bond_james"] = []string{"shaken, not stirred"}
	charInterests["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	charInterests["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	/*
		charInterests := map[string][]string{
			"bond_james":       {"shaken, not stirred"},
			"moneypenny_jenny": {"intelligence", "literature", "computer science"},
			"no_dr":            {"cats", "ice cream", "sunsets"},
		}
	*/

	for character, interests := range charInterests {
		fmt.Println(character)
		for interest, i := range interests {
			fmt.Println("\t", interest, i)
		}
	}

}
