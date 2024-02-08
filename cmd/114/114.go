package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not sirred"}
	mp := []string{"Miss", "Moneypenny", "I'm 008"}

	multislice := [][]string{jb, mp}

	for _, v := range multislice {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
}
