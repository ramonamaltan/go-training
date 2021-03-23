package main

import "fmt"

func main() {
	// making a bidirectional channel
	c := make(chan int)

	// send onto a channel
	go func() {
		c <- 42
	}()
	// receive from a channel
	fmt.Println(<-c)

	fmt.Println("---------")
	fmt.Printf("%T\n", c)
}
