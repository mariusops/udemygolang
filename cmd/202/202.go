/*
Method sets

Receivers Values
-----------------------------------------------
(t T) T and *T
(t *T) *T

The method sets of a type determines the interfaces that the type implements,
and the methods that can be called using a receiver of that type.

receiver type T works with both T and *T
receiver type *T works with only *T

https://go.dev/ref/spec#Method_sets




*/

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}

	//info(c) // cannot use c (type circle) as type shape in argument to info: circle does not implement shape (area method has pointer receiver)
	fmt.Println(c.area())
}
