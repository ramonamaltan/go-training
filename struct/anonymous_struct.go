package main

import "fmt"

// anonymous struct doesn't have a name
// no code pollution if person if never reused
func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1)
}
