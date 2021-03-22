package main

import (
	"fmt"
)

func main() {
	f := func(x int) {
		fmt.Println(x)
	}
	f(23)

	g := f
	g(78)
}
