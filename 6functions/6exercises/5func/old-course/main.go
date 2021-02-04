package main

import (
	"fmt"
)

func main() {

	// num, e := half(6)
	// fmt.Println(num, e)
	//modify half() to use a func expression e.g. to limit scope
	// half := func(n int) (int, bool) {
	// 	return n / 2, n%2 == 0
	// }
	// fmt.Println(half(2))

	fmt.Println(max(1, 2, 3))

	// fmt.Println(vals())
}

//half takes in an int and return the arg divided by 2 and true if even
func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

//write a func with one variadic parameter that finds the greatest num in a list of num
func max(nums ...int) int {
	var largest int
	for _, n := range nums {
		if n > largest {
			largest = n
		}
	}
	return largest

}

//what's the value expression of (t &&f) || (f && t) || !(f&&f) - false, false, true
func vals() bool {
	return (true && false) || (false && true) || !(false && false)
}
