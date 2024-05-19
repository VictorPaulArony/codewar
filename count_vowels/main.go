package main

import (
	"fmt"

	"count_vowels/vowels"
)

func main() {
	// Test cases
	fmt.Println(vowels.CountVowels("hello"))       // Output: 2
	fmt.Println(vowels.CountVowels("world"))       // Output: 1
	fmt.Println(vowels.CountVowels("programming")) // Output: 3
}
