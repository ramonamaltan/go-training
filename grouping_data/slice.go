package main

import "fmt"

func main() {
	// COMPOSITE LITERAL
	// slices have a variable length
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	// slice length
	fmt.Println(len(x))
	// accessing value by index position
	fmt.Println(x[4])

	// make([]type, length, capacity)
	y := make([]int, 10, 100)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	// ranging over slices
	for i, v := range x {
		fmt.Println(i, v)
	}
	// same as
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	// slicing a slice with : operator
	fmt.Println(x[1:3]) // [5, 7] not including last number
	fmt.Println(x[:3])  // [4, 5, 7]
}

// a SLICE allows you to group together VALUES of the same TYPE
