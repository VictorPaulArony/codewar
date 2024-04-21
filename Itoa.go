// We need a function that can transform a number (integer) into a string.

// What ways of achieving this do you know?

// Examples (input --> output):
// 123  --> "123"
// 999  --> "999"
// -100 --> "-100"
package main

import "fmt"

func intToString1(num int) string {
    return fmt.Sprintf("%d", num)
}

func main() {
    // Test case
    num := 123
    result := intToString1(num)
    fmt.Println(result) // Output: "123"
}
//**********************************************************************************************************************************//
package main

import (
    "fmt"
    "strconv"
)

func intToString2(num int) string {
    return strconv.Itoa(num)
}

func main() {
    // Test case
    num := 123
    result := intToString2(num)
    fmt.Println(result) // Output: "123"
}
