package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//Encode and Decode: i dont have to have to assign it to an var
	//directly to the wire
	//Decode: Reader
	//Encode: Writer

	fmt.Println("Hello, playground")
	fmt.Fprintln(os.Stdout, "Hello, playground")
	io.WriteString(os.Stdout, "Hello, playground")

}
