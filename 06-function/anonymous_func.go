package main

import "fmt"

func main() {
	foo()

	// anonymous or self executing functions
	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("Anonymous func with args ran:", x)
	}(42)
}

func foo() {
	fmt.Println("foo ran")
}
