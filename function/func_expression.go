package main

import "fmt"

func main() {
	// func expression: assign a function to a variable
	// functions are first class citizens
	f1 := func() {
		fmt.Println("my first func expression")
	}
	f1()

	f2 := func(x int) {
		fmt.Println("my first func expression with args:", x)
	}
	f2(52)
}
