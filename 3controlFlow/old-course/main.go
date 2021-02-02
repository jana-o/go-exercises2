package main

import "fmt"

func main() {
	//helloName()
	// printEvenNumbers()
	// fizzbuzz()
	sumMultiples3and5()

}

//program that asks user to enter name
func helloName() {
	name := ""
	fmt.Println("Please enter name: ")
	fmt.Scan(&name)
	fmt.Println("hello", name)
}

// print all even numbers betw 0-100
func printEvenNumbers() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func fizzbuzz() {
	for i := 1; i < 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println(i, "fizzbuzz")
		case i%3 == 0:
			fmt.Println(i, "fizz")
		case i%5 == 0:
			fmt.Println(i, "buzz")

		default:
			fmt.Println(i)
		}
	}
}

// all natural numbers below 10 that are multiples of 3 and 5 (=3,4,6,9) the sum = 23
// find all multiples of 3 and 5 below 1000
func sumMultiples3and5() {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
	//233.168
}

func new() {
	fmt.Println("new")
}
