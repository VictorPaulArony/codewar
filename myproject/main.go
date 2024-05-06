// main.go
package main

import (
	"fmt"
	"myproject/mylib" // Import your local package
)

func main() {
	mylib.MyFunction() // Use the function from your local package
	fmt.Println("Hello from my main program!")
}
