package main

import "fmt"

func main() {

	//Print every rune code point of the uppercase alphabet three times.

	for i := 65; i <= 90; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Printf("%v\t%U\t", i, i)
		}
		fmt.Printf("\n")
	}
	// birthyears()
	birthyears2()
	remainder5()
	remainder6()
	ifstate7()
	switch8()
}

func birthyears() {
	y := 1987
	for y <= 2018 {
		fmt.Println(y)
		y++
	}
}

func birthyears2() {
	y := 1987
	for {
		if y > 2018 {
			break
		}
		fmt.Println(y)
		y++
	}
}

func remainder5() {
	for i := 10; i < 100; i++ {
		if i%4 == 0 {
			fmt.Printf("%v\n", i)
		}
	}
}

func remainder6() {
	for i := 10; i <= 20; i++ {
		fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", i, i%4)
	}
}

func ifstate7() {
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			fmt.Printf("%v is even\n", j)
		} else {
			fmt.Printf("%v is uneven\n", j)
		}
	}
}

func switch8() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountain")
	case "swimming":
		fmt.Println("go to the sea")
	case "surfing":
		fmt.Println("go surfing")
	}
}
