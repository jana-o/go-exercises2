package main

import "fmt"

func main() {
	data := []float64{11, 22, 33}
	num := 4
	fmt.Println("the average is:", average(data...))
	fmt.Println("rec factorial", num, ":", recFactorial(num))
	fmt.Println("it factorial", num, ":", iterativeFactorial(num))
	fmt.Println("add all numbers from 1 to num:", simpleAdding(num))

}

func average(sl ...float64) float64 {
	total := 0.0
	for _, v := range sl {
		total += v
	}

	return total / float64(len(sl))

}

func recFactorial(num int) int {
	if num <= 1 {
		return 1
	}
	{
		return num * recFactorial(num-1)
	}
}

func iterativeFactorial(num int) int {
	result := 1
	if num <= 1 {
		return 1
	} else {
		for i := 1; i <= num; i++ {
			result *= i
		}
	}
	return result
}

func simpleAdding(num int) int {
	result := 0
	for i := 0; i <= num; i++ {
		result += i
	}
	return result
}

func simpleAdding2(num int) int {
	return (num/2)*(1+num) + (num%2)*((num/2)+1)
}
