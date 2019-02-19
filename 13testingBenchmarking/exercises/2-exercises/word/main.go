//Package word provides custom function for working with strings and words.
package word

import "strings"

//UseCount returns the number of times each word isused in a string.
// strings.Field breaks up string at white space, returns map, increment words
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

//Count returns the number of words in string.
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
