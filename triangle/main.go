package main

import (
	"fmt"

	"triangle/triangle"
)

func main() {
	fmt.Println(triangle.Triangle(3, 4, 5))   // true
	fmt.Println(triangle.Triangle(1, 2, 3))   // false
	fmt.Println(triangle.Triangle(2, 10, 12)) // false
}
