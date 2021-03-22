package main

import "fmt"

func main() {
	foo()
	bar("Ramona")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("ian", "flemming")
	fmt.Println(x)
	fmt.Println(y)

	// return a function from a function
	b2 := bar2()
	fmt.Printf("%T\n", b2) // func() int

	fmt.Println(b2()) // 451

	// all in one step
	fmt.Println(bar2()()) // 451
}

// func (r receiver -> attach to type) identifier(parameters) (return(s)) {...}
// basic func
func foo() {
	fmt.Println("hello from foo")
}

// takes in argument
func bar(s string) {
	fmt.Println("Hello,", s)
}

// return
func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

// multiple returns
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}

// return a function from a function
func bar2() func() int {
	return func() int {
		return 451
	}
}
