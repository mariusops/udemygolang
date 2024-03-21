// method sets part 1

package main

import "fmt"

type dog struct {
	name string
}

func (d dog) walk() {
	fmt.Println(d.name, "is walking")
}

func (d *dog) run() {
	d.name = "Fido"
	fmt.Println(d.name, "is running")
}

func main() {

	d1 := dog{
		name: "Henry",
	}
	d2 := &dog{
		name: "Padget",
	}

	fmt.Println(&d1)
	fmt.Println(&d2)

	d1.walk()
	d1.run()

	d2.walk()
	d2.run()

}
