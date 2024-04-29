// Write a function that checks if a given string (case insensitive) is a palindrome.

// A palindrome is a word, number, phrase, or other sequence of symbols that reads the same backwards as forwards, such as madam or racecar.
package main

import (
	"strings"
)

func IsPalindrome(str string) bool {
	// Convert the string to lowercase
	str = strings.ToLower(str)

	// Iterate over the string and compare characters from both ends
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		// If the characters at the current positions don't match, return false
		if str[i] != str[j] {
			return false
		}
	}
	// If the loop completes without finding any non-matching pairs, return true
	return true
}

func main() {
	// Test cases
	testCases := []string{"", "a", "radar", "Racecar", "Palindrome", "Madam", "12321", "hello"}

	for _, testCase := range testCases {
		if IsPalindrome(testCase) {
			println(testCase, "is a palindrome")
		} else {
			println(testCase, "is not a palindrome")
		}
	}
}
