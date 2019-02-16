package main

import "fmt"

var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)

	//code block inside code block
	//y is limited to that inner block
	{
		y := 42
		fmt.Println(y)
	}

	foo()
	fmt.Println(x)

	//incrementor
	incre()

}

func foo() {
	fmt.Println("helloo")
	x++
}

//closure: code block enclosing some variables
func incre() {
	a := incrementor()
	b := incrementor()
	// a and b have same identifier but diff scope, diff memory space where they are saved
	//keep scope narrow
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
