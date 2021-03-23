package main

import (
	"fmt"
	"math"
)

// methods set: set of methods attached to a type
// a non-pointer receiver: works with pointers or non-pointers
// a pointer receiver: only works with values that are pointers
// Receivers (t T) -> values T and *T
// Receivers (t T*) -> value *T
// IMPORTANT: the method set of a type determines the interfaces that the type implements

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{
		radius: 5,
	}
	info(&c)
	// info(c) doesn't run -> error: circle does not implement shape (area method has pointer receiver)
	// the method set of a type determines the interfaces that the type implements
}
