package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func printPerson(p person) {
	fmt.Println(p.first + " " + p.last + " is " + fmt.Sprint(p.age) + " years old.")
}

func changeFirst(p *person, newFirst string) {
	p.first = newFirst
}

func main() {

	/*p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}*/

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	sa2 := secretAgent{
		person: p2,
		ltk:    true,
	}

	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", p2, p2)

	printPerson(p2)

	changeFirst(&p2, "Penny")

	printPerson(p2)

	fmt.Println(sa1)
	fmt.Println(sa2)
}
