// altERnaTIng cAsE <=> ALTerNAtiNG CaSe
// Define String.prototype.toAlternatingCase (or a similar function/method such as to_alternating_case/toAlternatingCase/ToAlternatingCase in your selected language; see the initial solution for details) such that each lowercase letter becomes uppercase and each uppercase letter becomes lowercase. For example:

// ToAlternatingCase("hello world"); // => "HELLO WORLD"
// ToAlternatingCase("HELLO WORLD"); // => "hello world"
// ToAlternatingCase("hello WORLD"); // => "HELLO world"
// ToAlternatingCase("HeLLo WoRLD"); // => "hEllO wOrld"
// ToAlternativeCase("12345"); // => "12345" (Non-alphabetical characters are unaffected)
// ToAlternativeCase("1a2b3c4d5e"); // => "1A2B3C4D5E"
// ToAlternativeCase("String.prototype.toAlternatingCase"); // => "sTRING.PROTOTYPE.TOaLTERNATINGcASE"
// As usual, your function/method should be pure, i.e. it should not mutate the original string.

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
