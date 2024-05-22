// Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.

// invert([1,2,3,4,5]) == [-1,-2,-3,-4,-5]
// invert([1,-2,3,-4,5]) == [-1,2,-3,4,-5]
// invert([]) == []
// You can assume that all values are integers. Do not mutate the input array/list.
package main

//import "fmt"

func Invert(arr []int) []int {
    var ar []int
    for _, a := range arr {
        if a > 0 {
            ar = append(ar, -a)
        } else {
            ar = append(ar, -a)
        }
    }
    return ar
}

// func main() {
//     // Test cases
//     fmt.Println(Invert([]int{1, 2, 3, 4, 5}))    // Output: [-1, -2, -3, -4, -5]
//     fmt.Println(Invert([]int{1, -2, 3, -4, 5}))  // Output: [-1, 2, -3, 4, -5]
//     fmt.Println(Invert([]int{}))                 // Output: []
// }
