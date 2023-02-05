package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "alpha alpha alpha alpha alpha"
	str = replaceNth(str, "alpha", "beta", 3)
	fmt.Println(str)

	// case
	newString := "the quick brown FOX jumped over the lazy red DOG"
	fmt.Println(strings.ToLower(newString))
	fmt.Println(strings.ToUpper(newString))
	fmt.Println(strings.Title(newString))
	fmt.Println(strings.Title(strings.ToLower(newString)))
}

func replaceNth(s, old, new string, n int) string {
	// set an index to 0
	i := 0

	// go from 1 to no more than the nth position (passed to us)
	for j := 1; j <= n; j++ {
		// check to see if the current value of s, from
		// index i to the end, has the characters we want to replace
		x := strings.Index(s[i:], old)
		if x < 0 {
			// not found, so break out of the loop
			break
		}

		// found it, so add the index of the characters we are searching for
		// (old) to the current value of i
		i = i + x
		if j == n {
			// if the index of our loop equals the new value of i, then we have the start
			// of the characters we want to replace
			// we use the indexing function to return:
			// everything from the first position in the string to the current index of i
			// and conactenate the new word, then conactenate everything from after the
			// end of the old characters to the end of s to the end of this string
			// this is the replace function
			return s[:i] + new + s[i+len(old):]
		}

		// not ready to return yet, so increment i with the length of the old characters
		i += len(old)
	}
	return s
}
