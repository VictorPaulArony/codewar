package main

import (
	"fmt"

	"camel_case/camel"
)

func main() {
	fmt.Println(camel.CamelCase("hello case"))      // "HelloCase"
	fmt.Println(camel.CamelCase("camel case word")) // "CamelCaseWord"
}
