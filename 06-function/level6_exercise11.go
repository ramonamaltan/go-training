package main

import "fmt"

func main() {
	fmt.Println(rec(9))
}

func rec(n int) int {
	if n == 0 {
		return 1
	}
	return rec(n-1) * n
}
