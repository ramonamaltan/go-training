package main

import "fmt"

func main() {
	x := 54
	fmt.Println(&x)

	p1 := person{
		first:   "James",
		last:    "Bond",
		address: "Barnaby lane",
	}
	fmt.Println(p1)

	// p1 needs to be a pointer -> memory address
	changeMe(&p1)
	fmt.Println(p1)
}

type person struct {
	first   string
	last    string
	address string
}

// changeMe takes as arg a pointer to a type person
func changeMe(p *person) {
	p.address = "Blane"
}
