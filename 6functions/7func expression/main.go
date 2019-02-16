package main

import "fmt"

func main() {
	fmt.Println("Hello")

	//first class citizen => type just like the rest of it, can be passed anywhere
	f := func() {
		fmt.Println("my first func expression")
	}
	f()
	g := func(x int) {
		fmt.Println("year big brother", x)
	}
	g(1984)
}

//assign a function to a variable
