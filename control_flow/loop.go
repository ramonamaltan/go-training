package main

import "fmt"

func main() {
	// for statement with for ForClause
	// for init; condition; post { }
	for i := 0; i <= 10; i++ {
		fmt.Printf("the outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}

	// Modulo operator % (remainder)
	x := 43 % 40
	fmt.Println(x)

	// break keyword
	x = 0
	for {
		if x > 100 {
			break
		}
		if x%2 != 0 {
			fmt.Println(x)
		}
		x++
	}

	for y := 33; y <= 122; y++ {
		fmt.Printf("%v\t%b\t%#U\n", y, y, y)
	}
}
