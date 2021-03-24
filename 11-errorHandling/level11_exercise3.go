package main

import (
	"fmt"
)

type MyError struct {
	info string
}

func (m MyError) Error() string {
	return fmt.Sprintf("This is a custom error: %v", m.info)
}

func main() {
	m1 := MyError{
		info: "need more coffee!!!",
	}

	foo(m1)
}

func foo(e error) {
	fmt.Println("foo ran -", e)
}
