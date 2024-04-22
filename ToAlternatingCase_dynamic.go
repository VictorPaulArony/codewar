package main

import (
    "fmt"
)

func ToAlternatingCase(s string) string {
    // Convert the string to a slice of runes
    runes := []rune(s)
    
    // Iterate over each rune in the slice
    for i, char := range runes {
        // Check if the character is uppercase
        if char >= 'A' && char <= 'Z' {
            // Convert uppercase to lowercase
            runes[i] = char + 32
        } else if char >= 'a' && char <= 'z' {
            // Convert lowercase to uppercase
            runes[i] = char - 32
        }
    }
    
    // Convert the slice of runes back to a string and return it
    return string(runes)
}

func main() {
    // Test cases
    fmt.Println(ToAlternatingCase("hello world"))  // Output: "HELLO WORLD"
    fmt.Println(ToAlternatingCase("HELLO WORLD"))  // Output: "hello world"
    fmt.Println(ToAlternatingCase("hello WORLD"))  // Output: "HELLO world"
    fmt.Println(ToAlternatingCase("HeLLo WoRLD"))  // Output: "hEllO wOrld"
}
