package main

import (
	"fmt"
	"runtime"
	"sync"
)

//creating a race condition (see picture)
//use wait groups to wait

//https://godoc.org/sync#WaitGroup
//has package scope
//type WaitGroup= waits for a collection of goroutines to finish.
var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	//func (wg *WaitGroup) Add(delta int)
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	//func (wg *WaitGroup) Add(delta int)
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	//func (wg *WaitGroup) Done()
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
