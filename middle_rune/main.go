package main

import (
	"fmt"

	"middle_rune/mid"
)

func main() {
	// Test cases
	fmt.Println(mid.Middle("test"))    // Output: "es"
	fmt.Println(mid.Middle("testing")) // Output: "t"
	fmt.Println(mid.Middle("middle"))  // Output: "dd"
}
