package main

import "fmt"

func main() {
	s := "J"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#X\n", n)
	//0X4A => 4*16=64 then 74-64=10 => 0-1-A => 4A
}
