package main

import (
	"fmt"
)

var y = 42
var x = `backticks are strings`
var z string = "shaken not stirred"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
