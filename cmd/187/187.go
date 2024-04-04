// Concrete type vs interface type

// Concrete type is a type that you can directly instantiate or create a value from.

// Interface types are abstract. They represent behavior or type, but not a specific set of values.

package main

import "fmt"

// Concrete type
type Employee struct {
	Name string
	Age  int
}

// Interface type
type Reader interface {
	Read(p []byte) (n int, err error)
}

func main() {

	// Concrete type
	emp := Employee{Name: "John", Age: 30}
	fmt.Println(emp)

	// Interface type
	r := &Employee{Name: "Jane", Age: 25}
	r.Read([]byte("interface type"))

}

// Concrete type to satisfy the interface type
func (e *Employee) Read(p []byte) (n int, err error) {
	fmt.Printf("Reading data: %s\n", string(p))
	return 0, nil
}
