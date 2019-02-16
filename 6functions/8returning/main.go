package main

import "fmt"

func main() {

	x := bar()
	fmt.Printf("%T", x)
	i := x()
	fmt.Println(i)

	//fmt.Println(x())

	s1 := foo()
	fmt.Println(s1)

}

//return a func from a func that returns an int

func bar() func() int {
	return func() int {
		return 451
	}
}

func foo() string {
	s := "Hello world"
	return s
	//or simply
	//return "Hello world"
}
