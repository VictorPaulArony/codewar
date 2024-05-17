package main

import (
	"os"
	"strings"

	"ascii_art/ascii"
)

func main() {
	// to allow only 3 arguments to be used in the program
	if len(os.Args) != 3 {
		println("INVALID ARGUMENT FORMAT")
		os.Exit(0)
	}

	// read from the argument 2 that is the .txt file to output the ascii art
	data, err := os.ReadFile(os.Args[2])
	if err != nil {
		println("INVALID FILE")
		os.Exit(0)
	}
	// exclusion for the thinkertoy.txt to remove the courage return
	if os.Args[2] == "thinkertoy.txt" {
		lines := strings.Split(string(data), "\r\n")

		if len(lines) != 856 {
			println("THE TEXT FILE IS INCORRECT")
			os.Exit(1)
		}
		if len(os.Args) < 2 {
			println("Please provide text to display.")
			return
		}
		ascii.DisplayText(strings.Join(os.Args[1:2], ""), lines)
		// in the case of standard and shadoe text files
	} else {
		lines := strings.Split(string(data), "\n")
		ascii.DisplayText(strings.Join(os.Args[1:2], ""), lines)
	}
}
