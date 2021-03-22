package main

import "fmt"

func main() {
	sum(2, 7, 9, 4, 5, 3)
}

// variadic parameters ...
// variadic means 0 or more so sum() will also run
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are adding", v, "to the total which is now.", sum)
	}
	fmt.Println("The total is", sum)
	return sum
}
