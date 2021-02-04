package main

import (
	"fmt"
)

func main() {
	c := incrementor()
	sum := puller(c)
	for n := range sum {
		fmt.Println(n)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

//optional <- operator specifies the channel direction sebd ir receive
// if no direction goive -> chan is bidirectional
func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
