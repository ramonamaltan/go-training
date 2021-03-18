package main

import (
	"fmt"
)

var y = 42
var x = `backticks are strings`

// DECLARE a variable to be a certain TYPE
// and then ASSIGN a VALUE of that type to the variable
var z string = "shaken not stirred"

// create own type
type hotdog int

var a hotdog

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	// base 16
	fmt.Printf("%x\n", z)
	// binary number
	fmt.Printf("%b\n", y)

	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
