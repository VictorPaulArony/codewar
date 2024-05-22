package main

import (
	"fmt"
)

func main() {
	// Test cases
	fmt.Println(repeat.NonRepeatingLetter("stress")) // Output: "t"
	fmt.Println(repeat.NonRepeatingLetter("sTreSS")) // Output: "T"
	fmt.Printf(repeat.NonRepeatingLetter("ssttree")) // Output: ""
}
