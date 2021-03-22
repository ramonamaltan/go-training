package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{2938, 83, 0, 73, 930, 8, 27, 90, 8923, 28}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summer", "under", "gallantry"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
