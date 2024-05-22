package camel

import "strings"

func CamelCase(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}
