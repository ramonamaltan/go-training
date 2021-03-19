package main

import "fmt"

const (
	// every time shifting over 10 -> use iota (1, 2, 3)
	_ = iota
	// kb = 1024
	kb = 1 << (iota * 10)
	// mb = 1024 * kb
	mb = 1 << (iota * 10)
	// gb = 1024 * mb
	gb = 1 << (iota * 10)
)

func main() {
	a := 4
	fmt.Printf("%d\t\t%b\n", a, a) // 4 100

	// bit shifting
	b := a << 1
	fmt.Printf("%d\t\t%b\n", b, b) // 8 1000

	// every time shifting over 10
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
