package main

import "fmt"

//Pointer stores the memory address of another value located in computer memory
//A pointer references a location in memory (obtaining the value stores at that location is dereferencing)

func main() {
	a := 42
	fmt.Println(a)
	// & give me address &
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	// var b *int = &a => *int is a pointer to an int
	b := &a
	fmt.Println(b)
	//now dereference => give me the value stored at that address
	// * give me the value  *
	fmt.Println(*b)

	*b = 43
	fmt.Println(a)
	//b value changed and b is a => a changed

	//pointers are good if you have a huge piece of data
	//you only pass address around and not entire data
	//or if you de-reference value => change data ata location
}

//go has garbage collection
//Some built-in types, like maps and channels, are boxed
//(i.e. internally they are pointers to mutable structures), and are initialized using the make function
