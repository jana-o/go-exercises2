package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		log.Panicln(err)
		panic(err)
	}
}

// http://godoc.org/builtin#panic

func paniclnErr() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		log.Panicln(err)
		//		panic(err)
	}
}

/*
Panicln is equivalent to Println() followed by a call to panic().
*/

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
