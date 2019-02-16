package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//https://godoc.org/sync/atomic
//int64 => package atomic
//low-level atomic memory primitives useful for implementing synchronization algorithms.
//synchronization is better done with channels or the facilities of the sync package

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			//write to int64
			atomic.AddInt64(&counter, 1)
			//read int64
			fmt.Println("counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)
}

//go run -race main.go
