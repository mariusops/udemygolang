// defer
package main

import "fmt"

// use defer to close files, connections
// will close at the end of the main function in this case
func main() {
	defer foo()
	bar()
}

//func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
