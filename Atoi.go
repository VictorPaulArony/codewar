// Description
// We need a function that can transform a string into a number. What ways of achieving this do you know?

// Note: Don't worry, all inputs will be strings, and every string is a perfectly valid representation of an integral number.

// Examples
// "1234" --> 1234
// "605"  --> 605
// "1405" --> 1405
// "-7" --> -7
package main

import (
	"fmt"
	"strconv"
)

func stringToInt(str string) (int, error) {
	result, err := strconv.Atoi(str)
	if err != nil {
		return 0, err // Return an error if conversion fails
	}
	return result, nil // Return the converted integer and nil error
}

func main() {
	examples := []string{"1234", "605", "1405", "-7"}
	for _, example := range examples {
		number, err := stringToInt(example)
		if err != nil {
			fmt.Printf("Error converting '%s': %s\n", example, err)
		} else {
			fmt.Printf("'%s' --> %d\n", example, number)
		}
	}
}

