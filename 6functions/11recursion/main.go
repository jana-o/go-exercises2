package main

import "fmt"

//recursion == func calls itself
//can do it with loops as well
//not ery good for memory use

//factorial function
func main() {
	// fmt.Println(4 * 3 * 2 * 1)
	n := factorial(4)
	fmt.Println(n)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 4 * 3 * 2 * 1 *1
// return n * factorial(4-1) * factorial(3-1) * factorial(2-1) * 1

//easier way is to use a loop
func loop(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
