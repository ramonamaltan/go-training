package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println("Hello World")
	foo()
	fmt.Println("after the function I print this. This is the control flow.")
}

func foo() {
	fmt.Println("This is a function")
}