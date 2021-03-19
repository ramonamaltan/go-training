package main

import (
	"fmt"
)

const (
	// untyped constant
	x = 20
	// typed constant
	y int = 30
)

const (
	c = 2017 + iota
	d
	e
	f
)

func main() {
	num := 42
	// print number in decimal, binary and hex
	fmt.Printf("%d\t%b\t%#x\n", num, num, num)

	fmt.Println(x, y)

	a := 23
	fmt.Printf("%d\t%b\t%#X\n", a, a, a)

	// bit shifting
	b := a << 1
	fmt.Printf("%d\t%b\t%#X\n", b, b, b)

	fmt.Println(c, d, e, f)
}
