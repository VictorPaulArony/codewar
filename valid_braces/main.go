// Write a function that takes a string of braces, and determines if the order of the braces is valid. It should return true if the string is valid, and false if it's invalid.

// This Kata is similar to the Valid Parentheses Kata, but introduces new characters: brackets [], and curly braces {}. Thanks to @arnedag for the idea!

// All input strings will be nonempty, and will only consist of parentheses, brackets and curly braces: ()[]{}.

// What is considered Valid?
// A string of braces is considered valid if all braces are matched with the correct brace.

// Examples
// "(){}[]"   =>  True
// "([{}])"   =>  True
// "(}"       =>  False
// "[(])"     =>  False
// "[({})](]" =>  False
package main

import (
	"fmt"
	"strings"
)

func isValid(str string) bool {
	// 	stack := make([]rune, 0)

	// 	for _, char := range s {
	// 		if char == '(' || char == '[' || char == '{' {
	// 			stack = append(stack, char)
	// 		} else {
	// 			if len(stack) == 0 {
	// 				return false
	// 			}
	// 			top := stack[len(stack)-1]
	// 			if (char == ')' && top != '(') || (char == ']' && top != '[') || (char == '}' && top != '{') {
	// 				return false
	// 			}
	// 			stack = stack[:len(stack)-1]
	// 		}
	// 	}

	// 	return len(stack) == 0
	// }
	for i := 0; i < len(str); {
		a := len(str)
		str = strings.Replace(str, "()", "", -1)
		str = strings.Replace(str, "{}", "", -1)
		str = strings.Replace(str, "[]", "", -1)
		if a == len(str) {
			return false
		}
	}
	// fmt.Print(str)
	return true
}

func main() {
	fmt.Println(isValid("(){}[]"))   // true
	fmt.Println(isValid("([{}])"))   // true
	fmt.Println(isValid("(}"))       // false
	fmt.Println(isValid("[(])"))     // false
	fmt.Println(isValid("[({})](]")) // false
}
