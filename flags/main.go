package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		panic("INVALID ARGUMENT FORMAT")
	}
	if strings.HasSuffix(os.Args[2], ".txt") {
		err := os.WriteFile(os.Args[2], []byte(os.Args[1] + "\n"), 0o644)
		if err != nil {
			panic("INVALID")
		}
		fmt.Printf("This %s file has been created \n", os.Args[2])
	}
}
