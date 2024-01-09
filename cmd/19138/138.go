package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("Hello, my name is", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("Hello, I'm a secret agent named", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	sa1 := secretAgent{
		person: person{
			first: "Marius",
		},
		ltk: true,
	}

	//p1.speak()
	//sa1.speak()

	saySomething(p1)
	saySomething(sa1)

}
