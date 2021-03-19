package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	// delete from slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
