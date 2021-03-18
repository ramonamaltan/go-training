package main

import (
	"fmt"
)

var y = 42
var x = `backticks are strings`

// DECLARE a variable to be a certain TYPE
// and then ASSIGN a VALUE of that type to the variable
// Cannot reassign to another type
var z string = "shaken not stirred"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
