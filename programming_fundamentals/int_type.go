package main

import (
	"fmt"
)

var x int8

func main() {
	x = 127 // int8 can be between -128 and 127 (8 0s and 1s)
	y := 43.6824792
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
