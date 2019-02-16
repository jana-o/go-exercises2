package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   34,
	}

	//call the method from the value of type person
	p1.speak()

	//using defer
	defer foo()
	fmt.Println("Hello")
}

//attach a method to type person with
func (p person) speak() {
	fmt.Println("Hi, my name is:", p.first, "I'm", p.age, "years old.")
}

func foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}
