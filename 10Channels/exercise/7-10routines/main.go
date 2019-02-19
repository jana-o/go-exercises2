package main

import "fmt"

func main() {
	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")
}

//MY solution
// func main() {
// 	for i := 0; i < 10; i++ {
// 		c := make(chan int)
// 		go func() {
// 			for i := 0; i < 10; i++ {
// 				c <- i
// 			}
// 			close(c)
// 		}()
// 		for v := range c {
// 			fmt.Println(v)
// 		}
// 		fmt.Println("------")
// 	}
// 	fmt.Println("exit")
// }
