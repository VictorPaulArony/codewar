package main

import (
	"fmt"
	"os"
)

func canFormString(s1, s2 string) bool {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}
	return i == len(s1)
}

func main() {
	if len(os.Args) != 3 {
		return // Display nothing if the number of arguments is not 2
	}

	s1, s2 := os.Args[1], os.Args[2]
	if canFormString(s1, s2) {
		fmt.Println(s1) // Print the first string if it can be formed from the second string
	}
}
