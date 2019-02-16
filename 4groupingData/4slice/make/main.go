package main

import "fmt"

func main() {
	//x := []int{4, 5, 6, 8, 42}
	x := make([]int, 10, 100)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[3] = 10
	// x[10] = 200 x[10] does not work because only x[9]

	x = append(x, 200)
	fmt.Println((x))
	fmt.Println(len(x))

}
