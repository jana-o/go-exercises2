package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(counter("Helloooo"))
	fmt.Printf("%v\n", reverseArr("abc def"))
	fmt.Printf("%v\n", reverseStr("one two three"))

	longestWord("Hello hi holaaa")
	fmt.Println(timeConvert(12))

}

//create a vowel counter
func counter(str string) int {

	counter := 0

	for _, value := range str {
		switch value {
		case 'a', 'e', 'i', 'o', 'u':

			counter++
		}
	}
	return counter
}

//reverse array
func reverseArr(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

//reverse string
func reverseStr(str string) string {
	words := strings.Fields(str)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

//return the largest word in the string.
func longestWord(sen string) {

	words := strings.Split(sen, " ")

	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })

	fmt.Println("slice sorted by len:", words)
	fmt.Println("shortest word:", words[0])
	fmt.Println("longest words:", words[len(words)-1])
}

//time convert from 122min -> 2:2
func timeConvert(num int) string {

	return strconv.Itoa(num/60) + `:` + strconv.Itoa(num%60)

}
