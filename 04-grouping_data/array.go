package main

import "fmt"

func main() {
	// arrays have its capacity defined at creation time
	// once an array has allocated its size, the size can no longer be changed
	var x [5]int
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}
