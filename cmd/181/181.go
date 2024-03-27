package main

import "fmt"

type dude struct {
	first string
}

func renameValue(d dude, name string) dude {
	d.first = name
	return d
}

func renamePointer(d *dude, name string) {
	d.first = name
}

func main() {
	d1 := dude{
		first: "Henry",
	}
	fmt.Println(d1.first)

	d1 = renameValue(d1, "Harry")
	fmt.Println(d1.first)

	renamePointer(&d1, "Harold")
	fmt.Println(d1.first)

}
