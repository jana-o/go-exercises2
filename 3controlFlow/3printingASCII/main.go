package main

import "fmt"

//print out unicode character, loop over 122 and print out as unicode
func main() {
	for i := 0; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}
