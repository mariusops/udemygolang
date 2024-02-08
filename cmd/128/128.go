package main

import (
	"fmt"
)

type person struct {
	first           string
	last            string
	iceCreamFlavors []string
}

func printPerson(p person) {
	fmt.Println(p.first, p.last)
	for _, v := range p.iceCreamFlavors {
		fmt.Println(v)
	}
}

func main() {
	p1 := person{
		first:           "James",
		last:            "Bond",
		iceCreamFlavors: []string{"chocolate", "martini"},
	}

	p2 := person{
		first:           "Jenny",
		last:            "Moneypenny",
		iceCreamFlavors: []string{"strawberry", "vanilla"},
	}

	printPerson(p1)
	fmt.Println("\n")
	printPerson(p2)

}
