package main

import (
	"fmt"
	"regexp"
)

func longestWord(sen string) string {
	var stringsRegex = regexp.MustCompile(`\p{L}+`)
	var allStrings = stringsRegex.FindAllString(sen, -1)

	var longest = ``
	for _, element := range allStrings {
		if len([]rune(longest)) < len([]rune(element)) {
			longest = element
		}
	}

	return longest

}

func main() {

	fmt.Println(longestWord("sen"))

}
