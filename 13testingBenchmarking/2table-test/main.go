package main

import "fmt"

func main() {
	fmt.Println("2+3= ", mySum(2, 3))
	fmt.Println("2+3+1= ", mySum(2, 3, 1))
	fmt.Println("2+0= ", mySum(2))

}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
