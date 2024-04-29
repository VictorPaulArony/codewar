// Write a function that takes an array of strings as an argument and returns a sorted array containing the same strings, ordered from shortest to longest.

// For example, if this array were passed as an argument:

// ["Telescopes", "Glasses", "Eyes", "Monocles"]

// Your function would return the following array:

// ["Eyes", "Glasses", "Monocles", "Telescopes"]

// All of the strings in the array passed to your function will be different lengths, so you will not have to decide how to order multiple strings of the same length.



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
