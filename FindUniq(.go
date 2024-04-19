// There is an array with some numbers. All numbers are equal except for one. Try to find it!

// findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
// findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55
// It’s guaranteed that array contains at least 3 numbers.

// The tests contain some very huge arrays, so think about performance.

// This is the first kata in series:
package main

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

func main() {
    // Example input
    input := []float32{1.1, 2.2, 3.3, 2.2, 4.4, 5.5, 1.1, 3.3}

    // Find and print the first unmatched float32
    FindUniq(input)
}
