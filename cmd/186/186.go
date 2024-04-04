// Generics tutorial
// https://go.dev/doc/tutorial/generics

package main

import "fmt"

func main() {

	m4 := map[string]float64{"a": 1, "b": 2.2, "c": 3.3}
	SumIntsOrFloats(m4)

	// compare m4 map key a with b
	if m4["a"] > m4["b"] {
		fmt.Printf("%v > %v\n", m4["a"], m4["b"])
	} else {
		fmt.Printf("%v <= %v\n", m4["a"], m4["b"])
	}

}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
