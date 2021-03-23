package main

import "fmt"

type person struct {
	first string
	last  string
}

type human interface {
	speak() string
}

func (p *person) speak() string {
	return "I am " + p.first + " " + p.last
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func main() {
	p1 := person{
		first: "Ramona",
		last:  "Maltan",
	}

	// saySomething(p1)
	saySomething(&p1)
}
