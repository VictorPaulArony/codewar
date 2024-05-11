package main

import (
	"fmt"

	"add_two_arrays/arrays"
)

func main() {
	arr1 := []int{1, 1, 2, 3, 3, 4, 5, 5}
	arr2 := []int{7, 9, 10, 11, 12, 12}

	result := arrays.MergeArrays(arr1, arr2)
	fmt.Println("Result:", result) // Should print [1 2 3 4 5 7 9 10 11 12]
}
