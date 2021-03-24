package main

import "fmt"

// this will run
// buffers: channel allows n value(s) to sit in there

func main() {
	// buffered channels
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
}
