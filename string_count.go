// Given a string of digits, you should replace any digit below 5 with '0' and any digit 5 and above with '1'. Return the resulting string.

// Note: input will never be an empty string
package main

import (
	"fmt"
	"strings"
)

func replaceDigits(s string) string {
	count := ""

	// Iterate through each character in the string
	for _, word := range s {
		if word >= '5' {
			// If the digit is 5 or above, replace it with '1'
			count += "1"
		} else {
			// If the digit is below 5, replace it with '0'
			count += "0"
		}
	}

	return count
}

func main() {
	// Test cases
	fmt.Println(replaceDigits("123456")) // Output: 000111
	fmt.Println(replaceDigits("987654")) // Output: 111111
	fmt.Println(replaceDigits("012345")) // Output: 000111
}
