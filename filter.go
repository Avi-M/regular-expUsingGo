package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")

	// Find the first match.
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	// Find all matches and save to a slice, n less than 0 means return all matches, indicates length of slice if it's greater than 0.
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll", all)

	// Find index of first match, start and end position.
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	// Find index of all matches, the n does same job as above.
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	// Find first submatch and return array, the first element contains all elements, the second element contains the result of first (), the third element contains the result of second ().

	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	// Same as FindIndex().
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	// FindAllSubmatch, find all submatches.
	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchall)

	// FindAllSubmatchIndex,find index of all submatches.
	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)
}
