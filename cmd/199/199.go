package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (l ByLast) Len() int           { return len(l) }
func (l ByLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

/*
You can find type and method definitions in the sort package. https://pkg.go.dev/sort
Look at the example and copy this:

#####################

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

#####################

Now you can use ByAge as a type in your program. You can also change ByAge to ByLast, ByFirst, BySayings, etc. to sort by those fields.
You can also change the Less function to sort in the reverse order. For example, return a[i].Age > a[j].Age to sort by age in descending order.
*/

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	//fmt.Println(users)

	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, saying := range v.Sayings {
			fmt.Println("\t", saying)
		}
	}

	sort.Sort(ByAge(users))

	fmt.Println("Sorted by Age:\n")
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, saying := range v.Sayings {
			fmt.Println("\t", saying)
		}
	}

}
