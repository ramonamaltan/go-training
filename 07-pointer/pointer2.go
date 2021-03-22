package main

import "fmt"

// if you want to mutate a value pass in a value of an address instead of a value of an int

func main() {
	x := 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 50
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
