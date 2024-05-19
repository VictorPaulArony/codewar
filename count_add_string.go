// Given a non-negative integer, 3 for example, return a string with a murmur: "1 sheep...2 sheep...3 sheep...". Input will always be valid, i.e. no negative integers.
package main

import (
    //"strings"
    "strconv"
)

func countSheep(num int) string {
    res := ""
    for i := 1; i <= num; i++ {
        n := strconv.Itoa(i)
        res += n + " sheep..."
    }
    return res
}
