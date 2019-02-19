package main

import "fmt"

func main() {

	c := make(chan int)
	//send foo(c) or go func(){ put foo here and close}()
	go func() {
		for i := -0; i < 100; i++ {
			c <- i
		}
		close(c)
		//need to close range oer channel, otherwise deadlock
	}()

	//receive,not a go routine bc then it blocks
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

//send
// func foo(c chan<- int) {
// 	for i := -0; i < 100; i++ {

// 		c <- 42
// 	}
// 	close(c)
// }
