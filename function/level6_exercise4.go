package main

import "fmt"

func main() {
	p1 := person{
		first: "Ben",
		last:  "Howard",
		age:   33,
	}
	p1.speak()
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi I am", p.first, p.last, "and I am", p.age, "years old.")
}
