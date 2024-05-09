package main

import (
	"fmt"

	"stringtointslice/slice"
)

func main() {
	fmt.Println(slice.StringToIntSlice("A quick brown fox jumps over the lazy dog"))
	fmt.Println(slice.StringToIntSlice("Converted this string into an int"))
	fmt.Println(slice.StringToIntSlice("hello THERE"))
}
