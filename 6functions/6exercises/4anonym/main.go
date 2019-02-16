package main

import "fmt"

func main() {
	//Ex6: build and use an anonymous function

	func() {
		fmt.Println("Hello")
	}()
	fmt.Println("done")

	//Ex7: Assign a func to a variable, then call that func
	x := func() int {
		total := 0
		for i := 0; i < 10; i++ {
			total += i
		}
		return total
	}
	fmt.Println(x())

}

//ex7 solution
var x int
var g func()

func main2() {

	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	g = f
	g()
	fmt.Printf("this is g %T\n", g)

	fmt.Println("done")
}
