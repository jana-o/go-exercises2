package main

import (
	"fmt"
)

func main() {
	fmt.Println(bar2()())

}

func bar2() func() int {
	return func() int {
		return 451
	}
}
