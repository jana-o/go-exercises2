//Using goroutines, create an incrementer

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//solving race with mutex
//mutual exclusive lock

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	//go routines
	gs := 50
	m := sync.Mutex{}
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			// remove Go scheduler
			// runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(v)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value", incrementer)

	fmt.Println("------")
	main2()
}

func main2() {
	var wg sync.WaitGroup
	var incrementer int64

	gs := 50
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			//we increment here by 1
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value", incrementer)

}
