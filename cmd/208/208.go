/*
Hands-on exercise #2
This exercise will reinforce our understanding of method sets:
● create a type person struct
● attach a method speak to type person using a pointer receiver
	○ *person
● create a type human interface
	○ to implicitly implement the interface, a human must have the speak method
● create func “saySomething”
	○ have it take in a human as a parameter
	○ have it call the speak method
● show the following in your code
	○ you CAN pass a value of type *person into saySomething
	○ you CANNOT pass a value of type person into saySomething
● here is a hint if you need some help
	○ https://play.golang.org/p/FAwcQbNtMG

Receivers     Values
-----------------------------------------------
(t T) T  and  *T
(t *T)        *T

Why can't we pass a value of type person into saySomething?
...Because the speak method has a pointer receiver.

Cannot use p1 (variable of type person) as human value in argument to saySomething:
person does not implement human (method speak has pointer receiver)

First saySomething() is called, which calls speak() method.
speak() method has a pointer receiver, so it can only be called with a pointer to a person.
*/

package main

import (
	"fmt"
)

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello, I am", p.name)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "John",
	}

	// saySomething(p)
	saySomething(p1)

	// create a pointer to a person
	p2 := &person{
		name: "Jane",
	}

	saySomething(p2)
}
