package main

import "fmt"

// channels block
// this won't run

func main() {
	// making a channel
	c := make(chan int)

	// putting values on a channel
	c <- 42

	// taking values off a channel
	fmt.Println(<-c)
}
