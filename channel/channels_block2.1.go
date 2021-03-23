package main

import "fmt"

// this will run
// 2 go routines running

func main() {
	c := make(chan int)

	go func() {
		c <- 42 // takes value on
	}()

	fmt.Println(<-c) // takes value off
}
