package main

import (
	"fmt"
	"strings"
)

type person struct {
	first string
}

type superhero struct {
	p          person
	superpower string
}

func (s superhero) speak() {
	fmt.Println("Hello, I'm", s.p.first)
	fmt.Println("I got", strings.ToLower(s.superpower), "superpower\n")
}

func (p person) speak() {
	fmt.Println("Hello, I'm", p.first, "\n")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()

}

func main() {

	p1 := person{
		first: "Harry",
	}

	s1 := superhero{
		p: person{
			first: "Superman",
		},
		superpower: "Fly",
	}

	//p1.speak()
	//s1.speak()
	saySomething(p1)
	saySomething(s1)
}
