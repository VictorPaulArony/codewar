// Complete the solution so that it reverses all of the words within the string passed in.

// Words are separated by exactly one space and there are no leading or trailing spaces.

// Example(Input --> Output):

// "The greatest victory is that which requires no battle" --> "battle no requires which that is victory greatest The"
package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// Split the string into words
	words := strings.Fields(s)

	// Reverse the order of words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join the reversed words back with a single space
	result := strings.Join(words, " ")

	return result
}

func main() {
	input := "The greatest victory is that which requires no battle"
	output := reverseWords(input)
	fmt.Println(output)
}
