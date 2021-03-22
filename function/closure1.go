package main

import "fmt"

// closure: limiting scope of a variable
// we want to keep the scope as narrow as possible

// var x int
// var declared outside of func main can also be used in foo()

func main() {
	var x int // narrowed scope to func main
	fmt.Println(x)
	x++
	{
		y := 42
		fmt.Println(y)
		// y scope is this code block only
	}
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println("Hello")
}
