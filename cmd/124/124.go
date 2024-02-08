package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func printPerson(p person) {
	fmt.Println(p.first + " " + p.last + " is " + fmt.Sprint(p.age) + " years old.")
}

func changeFirst(p *person, newFirst string) {
	p.first = newFirst
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", p1, p1)

	printPerson(p1)

	changeFirst(&p1, "John")

	printPerson(p1)
}
