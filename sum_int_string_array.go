// Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.

// Return your answer as a number.
package main

import (
	"fmt"
	"strconv"
)

func SumMix(arr []interface{}) int {
	sum := 0
	for _, val := range arr {
		if num, ok := val.(int); ok {
			sum += num
		}
           if str, ok := val.(string); ok {
		  num, _ := strconv.Atoi(str) 
				sum += num
           }
	}
	return sum
}

func main() {
	arr := []interface{}{1, "2", "3", 9, "5", 6,"7"}
	fmt.Println(SumMix(arr))
}
