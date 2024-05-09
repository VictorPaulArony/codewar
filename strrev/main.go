package main

import (
	"fmt"

	"strrev/rev"
)

func main() {
	s := "Hello World!"
	s = rev.StrRev(s)
	fmt.Println(s)
}
