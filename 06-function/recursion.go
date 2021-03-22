package main

import "fmt"

// recursion: a function calls itself
// same can be done with loops

func main() {
	fmt.Println(factorial(8))
	fmt.Println(loop(8))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loop(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
