package main

import "fmt"

// SortByLength sorts a slice of strings by their lengths in ascending order
func SortByLength(arr []string) []string {
	n := len(arr)
	// Bubble sort algorithm
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if len(arr[j]) > len(arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	// Test the SortByLength function
	strings := []string{"apple", "banana", "cherry", "date", "fig"}
	fmt.Println("Original slice:", strings)
	sorted := SortByLength(strings)
	fmt.Println("Sorted by length:", sorted)
}
