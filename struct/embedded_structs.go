package main

import "fmt"

// create a value of type person
// struct allow to aggregate values of a type
type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "bond",
			age:   32,
		},
		ltk: true,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	// access struct
	fmt.Println(sa1.first, sa1.last, sa1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
