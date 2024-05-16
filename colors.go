package main

import (
	"flag"
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

func main() {
	// Define command-line flags
	colorPtr := flag.String("color", "", "Color to apply to the text")
	flag.Parse()

	// Validate the color flag
	if _, ok := colorMap[*colorPtr]; !ok {
		fmt.Println("Invalid color specified.")
		return
	}

	// Get the letters to be colored, if specified
	letters := ""
	args := flag.Args()
	if len(args) > 0 {
		letters = args[0]
	}

	// Get the input text
	text := ""
	if len(args) > 1 {
		text = strings.Join(args[1:], " ")
	}

	// Apply the color to the specified letters or the entire text
	coloredText := applyColor(text, letters, *colorPtr)

	// Print the colored text
	fmt.Println(coloredText)
}

// applyColor applies the specified color to the specified letters in the text
func applyColor(text, letters, color string) string {
	if text == "" {
		return ""
	}

	if letters == "" {
		// Color the entire text
		return fmt.Sprintf("%s%s%s", colorMap[color], text, colorMap["reset"])
	}

	// Color specific letters in the text
	var result strings.Builder
	for _, char := range text {
		if strings.ContainsRune(letters, char) {
			result.WriteString(fmt.Sprintf("%s%c%s", colorMap[color], char, colorMap["reset"]))
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}
