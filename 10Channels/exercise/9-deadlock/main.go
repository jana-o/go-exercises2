package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) //close so that range will eventually exit
	}()

	for n := range c { //use range here instead of for{}
		fmt.Println(n)
	}

	//	wDeadlock()
}

func wDeadlock() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		c <- i //write all 10 vals to chan
	}

	for {
		fmt.Println(<-c) //waiting for vals -> deadlock
	}

}
