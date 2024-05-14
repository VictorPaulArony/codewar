package main

import (
	"find_odds/find"
	"fmt"
)

func main() {
	input := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}
	res := find.FindOdd(input)
	fmt.Println(res)
}
