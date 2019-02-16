package main

import "fmt"

type person struct {
	name string
}

func main() {
	// x := 42
	//print the address of that value
	// fmt.Println(&x)

	//create a value of type person and print out value
	p1 := person{
		name: "jana",
	}
	fmt.Println(p1)

	//call changeMe()
	changeMe(&p1)
	fmt.Println(p1)
}

//create a func called “changeMe” with a *person as a parameter
//change the value stored at the *person address
func changeMe(p *person) {
	p.name = "nana"
	// (*p).name = "nana"

	//to dereference a struct, use (*value).field
	// p1.first
	// (*p1).first
}
