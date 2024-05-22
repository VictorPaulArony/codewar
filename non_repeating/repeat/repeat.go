package repeat

import (
	"strings"
)

func NonRepeatingLetter(words string) string {
	stack := make(map[rune]int)
	for _, word := range words {
		lower := strings.ToLower(string(word))
		stack[rune(lower[0])]++
	}
	for _, word := range words {
		lower := strings.ToLower(string(word))
		if stack[rune(lower[0])] == 1 {
			return string(word)
		}
	}
	return ""
}

// func NonRepeating(str string) string {
//     for _, c := range str {
//         if strings.Count(strings.ToLower(str), strings.ToLower(string(c))) < 2 {
// 	          return string(c)
// 	      }
//     }
//     return ""
// }
