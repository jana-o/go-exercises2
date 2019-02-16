package main

import "fmt"

//entfalten
//https://golang.org/ref/spec#Passing_arguments_to_..._parameters
func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	//take all the int in []int and put them in sum(xi...) in a slice
	//... for an unlimited number of ints => variadic
	fmt.Println("Total is", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index", i, "we are now adding", v, "to the total => now:", sum)
	}
	return sum
}
