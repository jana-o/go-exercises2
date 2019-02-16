package main

import "fmt"

func main() {

	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Println(a, b, c, d, e, f)
	fmt.Println("ex3")
	ex3()
	fmt.Println("ex4")
	ex4()
	ex5()
	iotayears()

}

func ex3() {
	const (
		//typed and untyped var
		x     = 42
		y int = 42
	)
	fmt.Println(x, y)
}
func ex4() {
	var a int
	a = 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	//shift b by 1
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}

func ex5() {
	fmt.Println("Hello 'string literal'")
}

func iotayears() {
	const (
		a = 2017 + iota
		b
		c
		d
	)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v", d)

}
