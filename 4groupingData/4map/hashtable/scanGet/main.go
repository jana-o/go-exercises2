package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

//get text from website, scan text and print out text
func main() {
	res, err := http.Get("https://www.w3.org/TR/PNG/iso_8859-1.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
