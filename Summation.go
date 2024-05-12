package main

import "fmt"

func summation(num int) int {
	count := 0

	// Iterate from 1 to num and count up each value
	for i := 1; i <= num; i++ {
		count += i
	}

	return count
}

// func Summation(n int) int {
//     return n * (n + 1) / 2
// }

// func Summation(n int) (sum int) {
// 	for i := 0; i <= n; i++ {
// 		sum += i
// 	}
// 	return
// }

func main() {
	// Test cases
	fmt.Println(summation(5)) // Output: 15 (1 + 2 + 3 + 4 + 5 = 15)
	fmt.Println(summation(8)) // Output: 36 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 = 36)
}
