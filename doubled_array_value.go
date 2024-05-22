// Given an array of integers, return a new array with each value doubled.

// For example:

// [1, 2, 3] --> [2, 4, 6]
package main

//import "fmt"

func DoubleArray(arr []int) []int {
    doubled := make([]int, len(arr))
    for i, num := range arr {
        doubled[i] = num * 2
    }
    return doubled
}

// func main() {
//     arr := []int{1, 2, 3, 4, 5}
//     doubled := doubleArray(arr)
//     fmt.Println(doubled) // Output: [2 4 6 8 10]
// }
