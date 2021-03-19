package main

import "fmt"

func main() {
	x := 40
	if x == 42 {
		fmt.Println("Our value was 42")
	} else if x == 40 {
		fmt.Println("Our value was 40")
	} else {
		fmt.Println("Our value was not 40 or 42")
	}
}
