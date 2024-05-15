package uniq

import "fmt"

func FindUniq(arr []float32) {
	match := make(map[float32]int)

	// Count the occurrences of each float32
	for _, ar := range arr {
		match[ar]++
	}

	// Output the unmatched float32s

	//  var uniq float32
	// for num, count := range counts {
	//     if count == 1 {
	//         uniq = num
	//         break // Found the unique number, so exit the loop
	//     }

	for _, ar := range arr {
		if match[ar] == 1 {
			fmt.Println(ar)
			return // Stop after printing the first unmatched number
		}
	}
}
