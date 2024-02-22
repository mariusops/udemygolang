package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("Hello, I'm", p.first)
}

func main() {

	p1 := person{
		first: "Harry",
	}

	p1.speak()
}
