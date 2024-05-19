package color

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
	colorPtr := flag.String("color", "", "Color to apply to the text")
	flag.Parse()

	color := *colorPtr
	if _, ok := colorMap[color]; !ok {
		fmt.Println("Invalid color specified.")
		return
	}

	letters := ""
	args := flag.Args()
	if len(args) > 0 {
		letters = args[0]
	}

	text := ""
	if len(args) > 1 {
		text = strings.Join(args[1:], " ")
	}

	coloredText := applyColor(text, letters, color)

	fmt.Println(coloredText)
	//return coloredText
}

func applyColor(text, letters, color string) string {
	if text == "" {
		return ""
	}

	if letters == "" {
		return fmt.Sprintf("%s%s%s", colorMap[color], text, colorMap["reset"])
	}

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
