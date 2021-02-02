package main

import "fmt"

func main() {

	n := hashBucket("Go", 12)
	fmt.Println(n) //here G => bucket 11

	printbuckets() //Ascii number, letter, bucket number
}

//returns the number of bucket is should go in
func hashBucket(word string, buckets int) int {
	//letter:= rune(word[0])
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}

//note from previous videos bc ASCII letter A starts w 65
//this shows us the bucket numbers (for 12 buckets)
func printbuckets() {
	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12)
	}
}
