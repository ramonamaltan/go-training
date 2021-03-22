package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age   int
}

type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	sort.SliceStable(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age", people)

	sort.Sort(ByName(people))
	fmt.Println("By name", people)

	// sort.SliceStable(people, func(i, j int) bool { return people[i].First < people[j].First })
	// fmt.Println("By name", people)
}
