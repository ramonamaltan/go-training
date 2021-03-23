package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("main")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(2)
	go func() {
		fmt.Println("Hello from one")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from two")
		wg.Done()
	}()

	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}
