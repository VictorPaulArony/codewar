// Write a function named first_non_repeating_letter† that takes a string input, and returns the first character that is not repeated anywhere in the string.

// For example, if given the input 'stress', the function should return 't', since the letter t only occurs once in the string, and occurs first in the string.

// As an added challenge, upper- and lowercase letters are considered the same character, but the function should return the correct case for the initial letter. For example, the input 'sTreSS' should return 'T'.

// If a string contains all repeating characters, it should return an empty string ("");

// † Note: the function is called firstNonRepeatingLetter for historical reasons, but your function should handle any Unicode character.
package main 
import ( 
  "fmt"
  "strings"
  )
func firstNonRepeatingLetter(words string) string {
  stack := make(map[rune]int)
  for _, word := words {
    lower := strings.ToLower(string(word))
    stack[rune(lower[0])]++
    for _, word := words {
      lower := strings.ToLower(string(word)) 
      if stack[rune(lower[0]) ==1 {
        return string(word)
        }
      }
    return ""
}
