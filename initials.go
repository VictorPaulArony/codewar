// Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

// The output should be two capital letters with a dot separating them.

// It should look like this:

// Sam Harris => S.H

// patrick feeney => P.F
package main

import (
	"fmt"
	"strings"
)

func initials(name string) string {
	// Split the name into first and last name
	parts := strings.Split(name, " ")
	// Extract the first letter of each part and convert to uppercase
	initials := strings.ToUpper(string(parts[0][0])) + "." + strings.ToUpper(string(parts[1][0]))
	return initials
}

func main() {
	fmt.Println(initials("Sam Harris"))       // Output: "S.H"
	fmt.Println(initials("patrick feeney"))   // Output: "P.F"
}
