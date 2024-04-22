// Implement a function that accepts 3 integer values a, b, c. The function should return true if a triangle can be built with the sides of given length and false in any other case.

// (In this case, all triangles must have surface greater than 0 to be accepted).

// Examples:

// Input -> Output
// 1,2,2 -> true
// 4,2,3 -> true
// 2,2,2 -> true
// 1,2,3 -> false
// -5,1,3 -> false
// 0,2,3 -> false
// 1,2,9 -> false 
package main

import "fmt"

func canFormTriangle(a, b, c int) bool {
    // Check if all sides are greater than zero
    if a <= 0 || b <= 0 || c <= 0 {
        return false
    }
    // Check the triangle inequality theorem
    return (a+b > c) && (a+c > b) && (b+c > a)
}

// func IsTriangle(a, b, c int) bool {
//   return a+b > c && b+c > a && a+c > b
// }

func main() {
    fmt.Println(canFormTriangle(3, 4, 5)) // true
    fmt.Println(canFormTriangle(1, 2, 3)) // false
    fmt.Println(canFormTriangle(2, 10, 12)) // false
}