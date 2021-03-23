package main

import "fmt"

func main() {
	// making a directional channel
	// send-only channel
	cs := make(chan<- int)

	// receive-only channel
	// cr := make(<-chan int)

	// send onto a channel
	go func() {
		cs <- 42
	}()

	// receive from a channel
	// fmt.Println(<-cr)

	fmt.Println("---------")
	fmt.Printf("%T\n", cs) // chan<- int
	// fmt.Printf("%T\n", cr) // <-chan int
}
