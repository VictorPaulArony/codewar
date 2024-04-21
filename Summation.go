// Summation
// Write a program that finds the summation of every number from 1 to num. The number will always be a positive integer greater than 0. Your function only needs to return the result, what is shown between parentheses in the example below is how you reach that result and it's not part of it, see the sample tests.

// For example (Input -> Output):

// 2 -> 3 (1 + 2)
// 8 -> 36 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8)
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
