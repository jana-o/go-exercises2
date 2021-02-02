package main

import "fmt"

func main() {
	// helloName()

	// fmt.Println("/// user name ///")
	// helloUser()

	inputRemainder()

	// fmt.Println("/// only even numbers ///")
	// onlyEven()

	// fmt.Println("/// FizzBuzz ///")
	// fizzBuzz()

	fmt.Println("/// sum of three and fives")
	threeFiveSum()
}
func helloName() {
	name := "jj"
	fmt.Println("Hello", name)
}

func helloUser() {
	var name string
	fmt.Println("Please enter name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}

func inputRemainder() {
	var inputA int
	var inputB int
	fmt.Println("Please enter number >10:")
	fmt.Scan(&inputA)
	fmt.Println("Please enter number <10:")
	fmt.Scan(&inputB)
	fmt.Println(inputA, "%", inputB, "=", inputA%inputB)
}
func onlyEven() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func fizzBuzz() {
	for num := 0; num < 20; num++ {
		if num%15 == 0 {
			fmt.Println(num, "FizzBuzz")
		} else if num%3 == 0 {
			fmt.Println(num, "Fizz")
		} else if num%5 == 0 {
			fmt.Println(num, "Buzz")
		} else {
			fmt.Println(num)
		}
	}
}

func threeFiveSum() {
	counter := 0
	for num := 0; num < 10; num++ {
		if num%3 == 0 {
			counter += num
		} else if num%5 == 0 {
			counter += num
		}
	}
	fmt.Println(counter)
}
