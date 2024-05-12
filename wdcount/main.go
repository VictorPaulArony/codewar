package main

import "fmt"

func WordMatch(str, random string) bool {
	c1 := 0
	c2 := 0
	l1 := len(str)
	l2 := len(random)
	for c1 < l1 && c2 < l2 {
		if str[c1] == random[c2] {
			c1++
		}
		c2++
	}
	return c1 == l1
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
