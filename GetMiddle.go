// You are going to be given a word. Your job is to return the middle character of the word. If the word's length is odd, return the middle character. If the word's length is even, return the middle 2 characters.

// #Examples:

// Kata.getMiddle("test") should return "es"

// Kata.getMiddle("testing") should return "t"

// Kata.getMiddle("middle") should return "dd"

// Kata.getMiddle("A") should return "A"
package main

import (
	"fmt"
)

func getMiddle(s string) string {
	length := len(s)

	// Check if the word's length is odd or even
	if length%2 == 0 { // Even length
		middle := length / 2
		return s[middle-1 : middle+1]
	} else { // Odd length
		middle := length / 2
		return string(s[middle])
	}
}

func main() {
	// Test cases
	fmt.Println(getMiddle("test"))   // Output: "es"
	fmt.Println(getMiddle("testing")) // Output: "t"
	fmt.Println(getMiddle("middle"))  // Output: "dd"
}
