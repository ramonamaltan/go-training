package main

import "fmt"

// use var for zero value
var y string

func main() {
	fmt.Println("printing the value of y", y)
	fmt.Printf("%T", y)
}
