package main

import (
	"fmt"
	"slice_add/slice"
)

func main() {
	fmt.Println(slice.SliceAdd([]int{1, 2, 3}, 4))
	fmt.Println(slice.SliceAdd([]int{}, 4))
}
