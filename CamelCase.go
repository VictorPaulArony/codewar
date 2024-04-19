// Write a method (or function, depending on the language) that converts a string to camelCase, that is, all words must have their first letter capitalized and spaces must be removed.

// Examples (input --> output):
// "hello case" --> "HelloCase"
// "camel case word" --> "CamelCaseWord"
package main 
import (
    "fmt"
    "strings"
    )
    func CamelCase(s string) string {
        words := strings.Split(s, " ")
        for i, word := range words {
            words[i] = strings.Title(word)
        }
        return strings.Join(words, "")
    }
    func main() {
	fmt.Println(CamelCase("hello case"))       // "HelloCase"
	fmt.Println(CamelCase("camel case word")) // "CamelCaseWord"
}
