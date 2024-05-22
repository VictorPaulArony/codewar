// Given a string s, your task is to return another string such that even-indexed and odd-indexed characters of s are grouped and the groups are space-separated. Even-indexed group comes as first, followed by a space, and then by the odd-indexed part.

// Examples
// input:    "CodeWars" => "CdWr oeas"
//
//	||||||||      |||| ||||
//
// indices:   01234567      0246 1357
// Even indices 0, 2, 4, 6, so we have "CdWr" as the first group.
// Odd indices are 1, 3, 5, 7, so the second group is "oeas".
// And the final string to return is "Cdwr oeas".
package main

import (
	"fmt"
	"strings"
)

func AlternatingSplit(str string) string {
	var s1 []string
	var s2 []string
	if str == "" {
		return ""
	}
	for i, s := range str {
		if i%2 == 0 {
			s1 = append(s1, string(s))
		} else {
			s2 = append(s2, string(s))
		}
	}
	str1 := strings.Join(s1, "")
	str2 := strings.Join(s2, "")
	return fmt.Sprintf("%v %v", str1, str2)
}

// func main() {
// 	input := "CodeWars"
// 	result := AlternatingSplit(input)
// 	fmt.Println(result) // Should print "CdWr oeas"
// }
