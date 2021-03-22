package main

import "fmt"

func main() {
	ii := []int{1, 7, 9, 34, 67}
	fmt.Println(foo(ii...))
	fmt.Println(bar(ii))
}

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
