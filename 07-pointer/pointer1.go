package main

import "fmt"

// Pointers point to locations in memory where values are stored (memory address)
// Ampersand & to get the memory address
// Asterisk * gives value stored at an address
// *int is an pointer to an int
// when to use pointers:
// - pass memory address where data is stored instead of the data itself (for large datasets)
// - change data at a location -> assign a new value to the address
// Everything in Go is PASS BY VALUE

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // *int = pointer to an int

	b := &a
	fmt.Println(b)  // memory address
	fmt.Println(*b) // dereference the address and gives value stored at an address

	*b = 43
	fmt.Println(a) // a becomes 43 because it's stored at the same memory address as b
}
