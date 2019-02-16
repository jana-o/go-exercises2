package main

import (
	"fmt"
)

//v1:
// func main() {
// 	x := 0
// 	foo(x)
// 	fmt.Println(x)
// }

// func foo(y int) {
// 	fmt.Println(y)
// 	y = 43
// 	fmt.Println(y)
// }

//v2 MUTATE
func main() {
	x := 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 43
	//de-reference, mutate value at address and everything pointing at it changes as well
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
