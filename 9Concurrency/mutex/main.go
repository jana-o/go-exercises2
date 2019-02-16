package main

import (
	"fmt"
	"runtime"
	"sync"
)

//A “mutex” is a mutual exclusion lock. ​
//Mutexes allow us to lock our code so that only one goroutine can access that locked chunk of code at a time.

//like check out book,no one else can use that variable until done using

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	//create mutex
	//Package sync: https://golang.org/pkg/sync/
	//type RWMutex: unlimited reads at same time but no write

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//		time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("counter:", counter)
}

//go run -race main.go
//use -race FLAG!!
