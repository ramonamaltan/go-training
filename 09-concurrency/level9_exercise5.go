package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var counter int64
	gr := 100
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("count", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end", counter)
}
