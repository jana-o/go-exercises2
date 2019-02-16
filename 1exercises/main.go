package main

import "fmt"

var xi = 42
var yi = "James Bond"
var z = true

type hotdog int

var x hotdog
var y int

func main() {
	exercise2()
	exercise4()
}

//print all of the VALUES to one single string. ASSIGN the
//returned value of TYPE string using the short declaration operator to a
// VARIABLE with the IDENTIFIER “s”
func exercise2() {
	s := fmt.Sprintf("%v\t%v\t%v", xi, yi, z)
	fmt.Println(s)
}

//1)create own type and assign value
//2) create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
//now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE

func exercise4() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
