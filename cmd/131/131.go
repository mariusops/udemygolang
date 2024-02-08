package main

func main() {

	person := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Money": 555,
			"Q":     777,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}

	println(person.first)
	println(person.friends["Money"])
	println(person.friends["Q"])
	println(person.favDrinks[0])
	println(person.favDrinks[1])

}
