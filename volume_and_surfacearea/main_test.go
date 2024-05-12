package main

import (
	"fmt"
	"testing"
)

func TestGetSize(t *testing.T) {
	tests := []struct {
		w, h, d  int
		expected [2]int
	}{
		{1, 1, 1, [2]int{6, 1}},   // Case 1
		{2, 2, 2, [2]int{24, 8}},  // Case 2
		{3, 4, 5, [2]int{94, 60}}, // Case 3
		// Add more test cases as needed
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			result := GetSize(test.w, test.h, test.d)
			if result != test.expected {
				t.Errorf("GetSize(%d, %d, %d) = %v; expected %v", test.w, test.h, test.d, result, test.expected)
			}
		})
	}
}
