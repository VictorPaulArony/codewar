package main

import (
	"fmt"
	"sort"
)

func SortIntegerTable(a []int) {
	sort.Ints(a)
}

func main() {
	s := []int{5, 4, 3, 6, 2, 1, 0}
	SortIntegerTable(s)
	fmt.Println(s)
}
