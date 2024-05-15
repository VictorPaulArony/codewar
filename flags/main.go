package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		panic("INVALID ARGUMENT FORMAT")
	}
	if strings.HasSuffix(os.Args[2], ".txt") {
		err := os.WriteFile(os.Args[2], []byte(os.Args[1]+"\n"), 0o644)
		if err != nil {
			panic("INVALID")
		}
		fmt.Printf("This %s file has been created \n", os.Args[2])
	} else {
		if strings.Contains(string(os.Args[2]), "--upper") {
			str, err := os.ReadFile(os.Args[1]) 
			if err != nil {
				panic("INVALID")
			}
			err = os.WriteFile(os.Args[1], []byte(strings.ToUpper(string(str))), 0o644)
			if err != nil {
				panic("INVALID")
			}
			fmt.Printf("This %s file has been modified \n", os.Args[2])
		} else if strings.Contains(string(os.Args[2]), "--lower") {
			str, err := os.ReadFile(os.Args[1]) 
			if err != nil {
				panic("INVALID")
			}
			err = os.WriteFile(os.Args[1], []byte(strings.ToLower(string(str))), 0o644)
			if err != nil {
				panic("INVALID")
			}
			fmt.Printf("This %s file has been modified \n", os.Args[2])
		} else if strings.Contains(string(os.Args[2]), "--cap") {
			str, err := os.ReadFile(os.Args[1]) 
			if err != nil {
				panic("INVALID")
			}
			err = os.WriteFile(os.Args[1], []byte(strings.Title(string(str))), 0o644)
			if err != nil {
				panic("INVALID")
			}
			fmt.Printf("This %s file has been modified \n", os.Args[2])
		}
	}
}
