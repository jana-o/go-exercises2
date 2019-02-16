package main

import "fmt"

func main() {
	for i := 0; i < +10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("the outer loop %v, the inner loop %d\n", i, j)
		}
	}
	fmt.Println("loopBreak:")
	loopBreak()
}

func loopBreak() {
	x := 0
	for {
		x++
		if x > 50 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
