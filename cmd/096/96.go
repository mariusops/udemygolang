package main

func main() {

	dessertArray := [...]string{"ice cream", "cake", "cookies", "brownies", "pie"}

	println(len(dessertArray))

	for _, v := range dessertArray {
		println(v)
	}

}
