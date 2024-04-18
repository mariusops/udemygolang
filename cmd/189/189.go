// JSON marshal

package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name string
	Age  int
}

func main() {

	emp1 := Employee{
		Name: "John",
		Age:  30,
	}

	emp2 := Employee{
		Name: "Jane",
		Age:  25,
	}

	employees := []Employee{emp1, emp2}

	bs, err := json.Marshal(employees)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))

}
