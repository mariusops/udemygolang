// Sort
package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type ByName []person
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func (b ByName) Len() int           { return len(b) }
func (b ByName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByName) Less(i, j int) bool { return b[i].first < b[j].first }

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.first, p.age)
}

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)

	sort.Sort(ByName(people))

	fmt.Println(people)

	sort.Sort(ByAge(people))

	fmt.Println(people)

}
