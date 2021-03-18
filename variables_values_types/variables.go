package main

import "fmt"

var y = 26

// declare there is a variable with the identifier z and type int
// assigns the zero value of type int to z
var z int

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a value (of a certain type)
	// cannot used outside function body
	x := 42
	fmt.Println(x)
	// variables declared with var can live outside function body
	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("print again", y)
}
