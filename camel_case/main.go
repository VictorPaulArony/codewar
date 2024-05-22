package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelCase("hello case"))      // "HelloCase"
	fmt.Println(CamelCase("camel case word")) // "CamelCaseWord"
}
