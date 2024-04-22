// Given two integers a and b, which can be positive or negative, find the sum of all the integers between and including them and return it. If the two numbers are equal return a or b.

// Note: a and b are not ordered!

// Examples (a, b) --> output (explanation)
// (1, 0) --> 1 (1 + 0 = 1)
// (1, 2) --> 3 (1 + 2 = 3)
// (0, 1) --> 1 (0 + 1 = 1)
// (1, 1) --> 1 (1 since both are same)
// (-1, 0) --> -1 (-1 + 0 = -1)
// (-1, 2) --> 2 (-1 + 0 + 1 + 2 = 2)
package main

import "fmt"
func GetSum(a, b int) int {
  count := 0
    if a > b {
      a,b = b,a
    }
  for i := a ; i <= b ; i++{
    count += i
  
  }
  return count
}
func main() {
    // Get the sum of numbers between 5 and -1
    sum := GetSum(5, -1)
    fmt.Println("Sum:", sum) // Output: Sum: 14
}
