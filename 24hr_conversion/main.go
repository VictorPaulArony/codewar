package main

import (
	"fmt"

	"24hr_conversion/converter"
)

func main() {
	arr := "06:13:13AM"
	maxSum := converter.TimeConversion(arr)
	fmt.Printf("%s\n", maxSum)
}
