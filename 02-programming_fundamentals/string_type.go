package main

import (
	"fmt"
)

func main() {
	s := "Hello"
	// or s := `Hello`
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// ASCII utf8
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		// binary
		fmt.Printf("%b\n", s[i])
		// hexadecimals
		fmt.Printf("%#X\n", s[i])
	}

	for i, v := range s {
		fmt.Println(i, v)
	}
}
