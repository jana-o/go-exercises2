package main

import "fmt"

var x int

type foo int

var y foo

//no int after const bar int => go will figure type out
// type aliases is not a good practice but it even works with int
const bar = 42

func main() {

	// y = 42
	y = bar
	fmt.Printf("%T\n", int(y))
	fmt.Printf("%T\n", bar)
	fmt.Println(bar)

}
