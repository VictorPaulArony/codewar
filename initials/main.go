package main

import (
	"fmt"

	"inititials/sample"
)

func main() {
	// va := init.Initial()
	// fmt.Println(va)
	fmt.Println(sample.Initial("Sam Harris"))     // Output: "S.H"
	fmt.Println(sample.Initial("patrick feeney")) // Output: "P.F"
}
