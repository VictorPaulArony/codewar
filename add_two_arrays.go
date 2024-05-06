package main

import (
	"fmt"
	"sort"
)

func MergeArrays(arr1, arr2 []int) []int {


	res := make([]int, len(arr1)+len(arr2))
	copy(res, arr1)
	copy(res[len(arr1):], arr2)

	sort.Ints(res) // Sort the resulting array

	// Remove duplicates from the sorted array
	var uniqueRes []int
	for i := 0; i < len(res); i++ {
		if i == 0 || res[i] != res[i-1] {
			uniqueRes = append(uniqueRes, res[i])
		}
	}

	return uniqueRes
}

func main() {
	arr1 := []int{1, 1, 2, 3, 3, 4, 5, 5}
	arr2 := []int{7, 9, 10, 11, 12, 12}

	result := MergeArrays(arr1, arr2)
	fmt.Println("Result:", result) // Should print [1 2 3 4 5 7 9 10 11 12]
}
