package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	a := 42
	b := "James Bond"
	c := true
	fmt.Println(a, b, c)
	// fmt.Println(x, y, z)
	// when not yet assigned prints zero values "" for strings, 0 for numeric types and false for bool
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
