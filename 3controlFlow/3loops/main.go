package main

import "fmt"

func main() {
	// loop1()
	// loop2()
	// loop3()
	// loopBreak()
	remainder()
}
func loop1() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("the outer loop %v, the inner loop %d\n", i, j)
		}
	}
	fmt.Println("loopBreak:")
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

func loop2() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")
}

func loop3() {
	x := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")
}

func remainder() {
	x := 1
	for {
		x++
		if x > 10 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
	fmt.Println("done.")
}
