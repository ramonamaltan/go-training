package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- person speak")
}

// any other type that has method speak() is also of type human
// interfaces allow a value to be of more than one type
// interface says: If you have this method you are also my type
// empty interface -> every value can be put in there
type human interface {
	speak()
}

func bar(h human) {
	// type assertion
	// asserting that it's type person or type secretAgent
	switch h.(type) {
	case person:
		fmt.Println("I am of type human and was passed into bar:", h.(person).first)
	case secretAgent:
		fmt.Println("I am of type human and was passed into bar:", h.(secretAgent).first)
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)
}
