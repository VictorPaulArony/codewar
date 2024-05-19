package main

import (
	"fmt"
	"os"
	"strings"

	"ascii_color/color"
)

func main() {
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	color.DisplayText(strings.Join(os.Args[3:4], ""), lines)
	
}
