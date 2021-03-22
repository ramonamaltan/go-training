package main

import "fmt"

func main() {
	switch { // missing switch expression defaults to true
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (2 == 2):
		fmt.Println("this should print")
		// fallthrough
		// if it's true then print the next one also
	case (4 == 4):
		fmt.Println("also true but doesn't print unless fallthrough")
	}

	// switching on a value
	switch "Bond" {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("james bond")
	default:
		fmt.Println("this is default")
	}

	// multiple cases
	n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Do No":
		fmt.Println("miss money")
	case "M":
		fmt.Println("M")
	default:
		fmt.Println("this is default")
	}
}
