package main

import (
	"fmt"

	"split_pairs/pairs"
)

func main() {
	input := "abcdefghd"
	pairs := pairs.Solution(input)
	fmt.Println(pairs)
}
