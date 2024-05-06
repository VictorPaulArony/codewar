// In helper.go
package ascii_art

import (
	"fmt"
	"strings"
)

func IsEnglish(word string) bool {
	// Your implementation to check if the word is English
	return true // Placeholder
}

func PrintWord(word string) {
	var res []string
	for _, char := range word {
		first := fmt.Sprintf("%d", (int(char)-32)*9-2)
		res = append(res, first)
	}
	fmt.Println(strings.Join(res, "\n"))
}
