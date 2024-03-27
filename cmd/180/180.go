package main

import "fmt"

type dog struct {
	first string
}

type youngin interface {
	walk()
	run()
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I am running")
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I am walking")
}

func youngRun(y youngin) {
	y.run()
	y.walk()
}

func main() {
	d1 := dog{"Henry"}
	youngRun(d1) // doesnt run

	d2 := dog{"Padget"}
	youngRun(d2)
}
