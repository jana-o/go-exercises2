//Using goroutines, create an incrementer

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	incrementer := 0
	//go routines
	gs := 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value", incrementer)

}
