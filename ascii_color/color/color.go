package color

import (
	"fmt"
	"strings"
)

// Define ANSI escape codes for colors
var colorMap = map[string]string{
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",
	"white":  "\033[37m",
	"reset":  "\033[0m", // Reset color
}

// DisplayText displays the provided text along with content lines
func DisplayText(input string, contentLines []string, color string, letters string) {
	if input == "" {
		return
	}

	// Make newline and tab printable in the terminal output
	input = strings.ReplaceAll(input, "\n", "\\n")
	input = strings.ReplaceAll(input, "\t", "\\t")

	wordSlice := strings.Split(input, "\\n")

	for _, word := range wordSlice {
		if word == "" {
			fmt.Print("\n")
		} else {
			if IsEnglish(word) {
				PrintWord(word, contentLines, color, letters)
			} else {
				fmt.Print("Invalid input: not accepted")
			}
		}
	}
}

// IsEnglish checks if a word contains only printable ASCII characters
func IsEnglish(word string) bool {
	for _, char := range word {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}

// PrintWord prints a word if it exists in the content lines
func PrintWord(word string, contentLines []string, color string, letters string) {
	linesOfSlice := make([]string, 1)

	for _, v := range word {
		for i := 0; i < 2; i++ {
			charLine := contentLines[v]
			if letters == "" || strings.ContainsRune(letters, v) {
				linesOfSlice[i] += ApplyColor(charLine, letters, color)
			} else {
				linesOfSlice[i+1] += charLine
			}
		}
	}

	fmt.Print(strings.Join(linesOfSlice, " "))
}

// ApplyColor applies the specified color to the text
func ApplyColor(text, letters, color string) string {
	colorCode, exists := colorMap[color]
	if !exists {
		fmt.Println("Invalid color specified.")
		return text
	}
	return fmt.Sprintf("%s%s%s", colorCode, text, colorMap["reset"])
}
