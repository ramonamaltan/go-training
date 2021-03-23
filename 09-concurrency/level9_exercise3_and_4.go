package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	counter := 0
	gr := 100
	wg.Add(gr)

	var mu sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println("count", counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end", counter)
}
