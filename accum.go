package main

import (
	"fmt"
	"strings"
)

func Accum(words string) string {
	// the space where to store the final result
	result := ""
	for i, word := range words {
		// change all the given chararcters into uppper case form first
		up := strings.ToUpper(string(word))
		// itterate through all the strings up to i
		rep := strings.Repeat(up, i)
		// change the itterated words to lower case
		low := strings.ToLower(rep)
		// concatination of the lower case and the upper case
		result += up + low
		// adding of the - to all except to the n-1 word
		if i != len(words)-1 {
			result += "-"
		}
	}
	return result
}

// func Accum(s string) string {
//     sl := strings.Split(s,"")
//     for i, letter := range sl  {
//        sl[i] = strings.Title(strings.Repeat(strings.ToLower(letter), i + 1))
//     }
//     return strings.Join(sl, "-")
// }

func main() {
	testCases := []string{"abcd", "RqaEzty", "cwAt"}
	for _, testCase := range testCases {
		fmt.Printf("accum(\"%s\") -> \"%s\"\n", testCase, Accum(testCase))
	}
}

// This time no story, no theory. The examples below show you how to write function accum:

// Examples:
// accum("abcd") -> "A-Bb-Ccc-Dddd"
// accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
// accum("cwAt") -> "C-Ww-Aaa-Tttt"
// The parameter of accum is a string which includes only letters from a..z and A..Z.
