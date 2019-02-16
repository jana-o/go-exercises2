package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

//channels BLOCK!
//code does not work bec nothing to start c <-42
func notWorking() {
	c := make(chan int)
	c <- 42

	fmt.Println(<-c)
}

//buffer channel => better to stay away from buffer
//allows 1 to sit in there
func bufferChannel() {
	c := make(chan int, 1)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

//unsuccessful buffer
func unsuccessfulbuffer() {
	c := make(chan int, 1)

	go func() {
		c <- 42
		c <- 43
	}()

	fmt.Println(<-c)
	//would work if
	//c := make(chan int, 2)
	//fmt.Println(<-c)
}
