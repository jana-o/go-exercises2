package main

import (
	"exercises2/13testingBenchmarking/exercises/2-exercises/quote"
	"exercises2/13testingBenchmarking/exercises/2-exercises/word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
