package main

import "fmt"

func main() {
	//%d decimal \t => tab %b= binary
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	//bit shifting y equal x but shifted by 1
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	iotaf()
}

func iotaf() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}

//iota and bit shifting
// package main
// import "fmt"

// const (
// 	_  = iota             // 0
// 	KB = 1 << (iota * 10) // 1 << (1 * 10)
// 	MB = 1 << (iota * 10) // 1 << (2 * 10)
// 	GB = 1 << (iota * 10) // 1 << (3 * 10)
// 	TB = 1 << (iota * 10) // 1 << (4 * 10)
// )

// func main() {
// 	fmt.Println("binary\t\t\t\tdecimal")
// 	fmt.Printf("%b\t\t\t", KB)
// 	fmt.Printf("%d\n", KB)
// 	fmt.Printf("%b\t\t", MB)
// 	fmt.Printf("%d\n", MB)
// 	fmt.Printf("%b\t", GB)
// 	fmt.Printf("%d\n", GB)
// }
