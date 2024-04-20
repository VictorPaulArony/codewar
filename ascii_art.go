package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Read ASCII art from the "sample.txt" file
	asciiData, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Error reading sample.txt:", err)
		return
	}

	// Split the ASCII data into individual words
	words := strings.Split(string(asciiData), "\n\n\n\n")

	// Display the ASCII art for each word
	for _, word := range words {
		fmt.Println(word)
		fmt.Println()
	}

	// Check if there are no words
	if len(words) == 0 {
		fmt.Println("No ASCII art found in sample.txt")
	}
}
