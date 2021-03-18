package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	foo()
	fmt.Println("after the function I print this. This is the control flow.")
}

func foo() {
	fmt.Println("This is a function")
}
