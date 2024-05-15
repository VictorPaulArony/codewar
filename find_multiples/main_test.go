package main

import (
	"reflect"
	"testing"
)

func TestFindMultiples(t *testing.T) {
	// Test cases
	testCases := []struct {
		integer, limit int
		expected       []int
	}{
		{2, 10, []int{2, 4, 6, 8, 10}},
		{3, 15, []int{3, 6, 9, 12, 15}},
		{5, 25, []int{5, 10, 15, 20, 25}},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		// Call the function under test
		result := FindMultiples(tc.integer, tc.limit)

		// Compare the result with the expected value
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("FindMultiples(%d, %d) = %v; expected %v", tc.integer, tc.limit, result, tc.expected)
		}
	}
}
