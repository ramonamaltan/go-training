package main

import (
	"fmt"
	"runtime"
	"sync"
)

// locking access to a variable/ a certain chunk of code
// prevents race conditions -> counter 100

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched() // allow sth else to run (like sleep)
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	fmt.Println("count", counter)
}
