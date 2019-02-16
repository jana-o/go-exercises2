package main

import "fmt"

func main() {
	s := "Hello, Jana"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	//in decimal:
	//[72 101 108 108 111 44 32 74 97 110 97] => 32 space incl

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	// utf-8 code point:
	//' U+004A 'J' U+0061 'a' U+006E 'n' U+0061 'a'

	for i, v := range s {
		//fmt.Println(i, v)
		//hexadecimal
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
