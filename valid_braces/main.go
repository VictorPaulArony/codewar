
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
