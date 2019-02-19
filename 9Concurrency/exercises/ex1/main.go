package main

import (
	"fmt"
	"runtime"
	"sync"
)

//launch two additional go routines
//each additional go routine should print something out
//use waitgroups to make sure each go finishes before program exits
func main() {

	fmt.Println("begin CPU", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	//my variable wg is of type Waitgroup from package sync
	wg.Add(2)

	go func() {
		fmt.Println("Hello from 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello from 2")
		wg.Done()

	}()
	fmt.Println("mid CPU", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end CPU", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())

}
