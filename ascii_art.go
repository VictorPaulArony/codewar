package main

import (
	"fmt"
	"os"
	"strings"
)

// DisplayText displays the provided text along with content lines
func DisplayText(input string, contentLines []string) {
	if input == "" {
		return
	}

	if input == "\\n" || input == "\n" {
		fmt.Println()
		return
	}

	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\\t", "\t")

	wordslice := strings.Split(input, "\\n")

	for _, word := range wordslice {
		if word == "" {
			fmt.Println()
		} else {
			if IsEnglish(word) {
				PrintWord(word, contentLines)
			} else {
				fmt.Println("Invalid input:", word)
				// Optionally continue processing other words
			}
		}
	}
}

// IsEnglish checks if a word contains only English alphabets
func IsEnglish(word string) bool {
	for _, char := range word {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') {
			return false
		}
	}
	return true
}

// PrintWord prints a word if it exists in the content lines
func PrintWord(word string, contentLines []string) {
	linesOfSlice := make([]string, 9)

	for _, v := range word {
		for i := 1; i <= 9; i++ {
			linesOfSlice[i-1] += contentLines[int(v-32)*9+i]
		}
	}

	fmt.Print(strings.Join(linesOfSlice, "\n"))
}

func main() {
	// Assuming you want to read from a file named "standard.txt"
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	contentLines := strings.Split(string(content), "\n")

	if len(contentLines) != 856 {
		fmt.Println("The content does not have 856 lines.")
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Please provide text to display.")
		return
	}

	DisplayText(strings.Join(os.Args[1:], ""), contentLines)
}
