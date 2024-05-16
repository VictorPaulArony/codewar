package main

import (
	"fmt"

	"max_and_min/min_max"
)

func main() {
	arr := []int{3, 5, 10, 9, 2, 7}
	min, max := min_max.MinMax(arr)
	fmt.Printf("Min: %d, Max: %d\n", min, max)
}
