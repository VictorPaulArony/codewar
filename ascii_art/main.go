package main

import (
	"fmt"
	"os"
	"strings"
)

func display(str string) []string {
	words := strings.Split(str, "\n")
	for _, word := range words {
		if word == "" {
			FMT.Println()
		} else {
			if IsEnglish(word) {
				printword(word, res)
			} else {
				fmt.Println("INVALID INPUT: ", word)
			}
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("No input text to display!")
		return
	}
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("INVALID TEXT FILE")
		return
	}
}
