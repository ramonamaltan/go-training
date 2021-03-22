package main

import (
	"fmt"
	"math"
)

func main() {
	s1 := square{
		length: 10.9348,
	}

	c1 := circle{
		radius: 12.945,
	}

	info(c1)
	info(s1)
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
