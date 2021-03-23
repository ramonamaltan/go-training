package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// use package atomic to do same as mutex
// prevent race

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("GoRoutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	fmt.Println("count", counter)
}
