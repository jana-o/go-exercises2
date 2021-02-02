package main

import (
	"fmt"
	"time"
)

func main() {
	//We’re executing an external call that returns its result on a channel c1 after 2s.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

//select implementing a timeout. res := <-c1 awaits the result
//and <-Time.After awaits a value to be sent after the timeout of 1s.
//Since select proceeds with the first receive that’s ready,
//we’ll take the timeout case if the operation takes more than the allowed 1s

//Output: the first operation timing out and the second succeeding.
