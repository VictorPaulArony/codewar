package main

import "fmt"

func WordMatch(str, random string) bool {
	if str == "" || random == "" {
		return false
	}

	c1 := make(map[rune]int)
	c2 := make(map[rune]int)
	for _, char := range str {
		c1[char]++
	}
	for _, char := range random {
		c2[char]++
	}
	for char, count := range c1 {
		c2CharCount, ok := c2[char]
		if !ok || count > c2CharCount {
			return false
		}
	}
	return true
}

func main() {
	tests := [][]string{
		{"123", "123"},
		{"faya", "fgvvfdxcacpolhyghbreda"},
		{"faya", "fgvvfdxcacpolhyghbred"},
		{"error", "rrerrrfiiljdfxjyuifrrvcoojh"},
		{"quarante deux", "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "},
		{"", ""},
	}

	for _, test := range tests {
		if len(test) != 2 {
			continue
		}
		input, chars := test[0], test[1]
		if WordMatch(input, chars) {
			fmt.Println(input)
		}
	}
}
