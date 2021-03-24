package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	// receive
	// range loops pulls off of the channel until the channel is closed
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
