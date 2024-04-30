package main
import (
    "strings"
   // "unicode"
    )
func IsPangram(input string) bool {
    input = strings.ToLower(input)
    letterCount := make(map[rune]bool)
    for _, char := range input {
        if 'a' <= char && char <= 'z' {
            letterCount[char] = true
        }
    }
    for i := 'a'; i <= 'z'; i++ {
        if !letterCount[i] {
            return false
        }
    }
    
    return true
}
