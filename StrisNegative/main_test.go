package main

import "testing"

func StrIsNegative(t *testing.T, input string) string {
    if input[0] == '-' {
        return "N"
    } else {
        return "P"
    }
}

func StrIsNegativeTest(t *testing.T) {
    testCases := []struct {
        input    string
        expected string
    }{
        {"585", "P"},      // Positive number
        {"-58", "N"},      // Negative number
        {"55s44", "!"},    // Invalid characters
        {"101-1331", "!"}, // Invalid characters
        {"5544-", "!"},    // Invalid characters
    }

    for _, tc := range testCases {
        t.Run(tc.input, func(t *testing.T) {
            actual := StrIsNegative(t, tc.input)
            if actual != tc.expected {
                t.Errorf("StrIsNegative(%s): expected %s, got %s", tc.input, tc.expected, actual)
            }
        })
    }
}