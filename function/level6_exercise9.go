package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5}
	fmt.Println(foo(sum, ii))
}

func sum(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
