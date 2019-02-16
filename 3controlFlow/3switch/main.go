package main

import "fmt"

func main() {
	//switch bool
	switch {
	case false:
		fmt.Println("should not print")
	case (3 == 3):
		fmt.Println("should print 1")
		fallthrough
		//next case also gets printed with fall through
	case (4 == 3):
		fmt.Println("should not print")
		fallthrough
	case (4 == 4):
		fmt.Println("should print 2")
		fallthrough
	default:
		fmt.Println("this is default")
	}

	switchValue()
}

func switchValue() {
	n := "Bond"
	switch n {
	case "Moneypenny":
		fmt.Println("should not print")
	case "Bond":
		fmt.Println("should print Bond")
	default:
		fmt.Println("this is default")
	}
}
