package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var answer1, answer2 string

	fmt.Print("Fav number: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav color: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2)

	example2()
	example3()
}

//create a file and save as names.txt
func example2() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("James Bond")

	io.Copy(f, r)

}

//open a file
func example3() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("could not open file", f)
		return
	}
	defer f.Close()

	//needs bs[]
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}
