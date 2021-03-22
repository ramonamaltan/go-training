package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	// append to slice
	x = append(x, 6, 9, 0, 20830)
	fmt.Println(x)

	y := []int{67, 97, 0, 8, 4646, 29}
	x = append(x, y...)
	fmt.Println(x)
}
