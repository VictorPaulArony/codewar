// Count the number of Duplicates
// Write a function that will return the count of distinct case-insensitive alphabetic characters and numeric digits that occur more than once in the input string. The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits.

// Example
// "abcde" -> 0 # no characters repeats more than once
// "aabbcde" -> 2 # 'a' and 'b'
// "aabBcde" -> 2 # 'a' occurs twice and 'b' twice (`b` and `B`)
// "indivisibility" -> 1 # 'i' occurs six times
// "Indivisibilities" -> 2 # 'i' occurs seven times and 's' occurs twice
// "aA11" -> 2 # 'a' and '1'
// "ABBA" -> 2 # 'A' and 'B' each occur twice
package main

import (
    "fmt"
    "strings"
)

func CountDuplicates(input string) int {
    // Normalize the input to lowercase to handle case insensitivity
    input = strings.ToLower(input)
    
    // Create a map to count occurrences of each character
    charCount := make(map[rune]int)
    
    // Iterate over each character in the string and count occurrences
    for _, char := range input {
        charCount[char]++
    }
    
    // Count characters that appear more than once
    duplicateCount := 0
    for _, count := range charCount {
        if count > 1 {
            duplicateCount++
        }
    }
    
    return duplicateCount
}

func main() {
    examples := []string{"abcde", "aabbcde", "aabBcde", "indivisibility", "Indivisibilities", "aA11", "ABBA"}
    for _, example := range examples {
        fmt.Printf("'%s' -> %d\n", example, CountDuplicates(example))
    }
}
