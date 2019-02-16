package main

import "fmt"

func main() {
	//ex 8 Create a func which returns a func
	//assign the returned func to a variable

	f := foo1()
	fmt.Println(f())

	//callback
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}
	x := foo(g, []int{1, 2, 3, 4, 5})
	fmt.Println(x)
}

func foo1() func() int {
	return func() int {
		return 42
	}
}

//ex 9 CALLBACKS
// pass a func into a func as an argument

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
