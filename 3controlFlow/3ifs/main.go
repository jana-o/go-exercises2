package main

import "fmt"

func main() {
	if x := 42; x == 42 {
		fmt.Println("001")
	}
	//fmt.Println outside if does not work here
	// fmt.Println(x)
	xscope()
	ifelse()
}

func xscope() {
	x := 42
	if x == 42 {
		fmt.Println("001")
	}
	//here it workds because var declared outside => scope here
	fmt.Println(x)

}

func ifelse() {
	y := 42
	if y == 40 {
		fmt.Println("our value was 40")
	} else if y == 41 {
		fmt.Println("our value was 41")
	} else {
		fmt.Println("our value was not 40, 41")
	}
}
