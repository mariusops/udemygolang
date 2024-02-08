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
	fmt.Println("\n")

	pMap := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(pMap)

	for _, v := range pMap {

		fmt.Println(v.first)
		fmt.Println(v.last)
		for _, val := range v.iceCreamFlavors {
			fmt.Println(val)
		}
		fmt.Println("-------")
	}

}
