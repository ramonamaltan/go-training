package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("defer foo")
}

func bar() {
	fmt.Println("run bar")
}
