package main

import (
	"fmt"
)

type customErr struct {
	info string
}

//any of value type customErr will implicitly implement Error interface
//bec it has method Error() attached to it

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more coffee",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran -", e, "\n")

	//	fmt.Println("foo ran -", e, "\n", e.info) => e.info does not work bc info is attached to customErr not e
	//would need assertion for interface
	//e.(customErr).info
	//if you only have 1 value x=y  then you use conversion

}

type hotdog int

func mainConversionExample() {
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
