package negative

import (
	"fmt"
	"strconv"
)

func StrIsNegative(str string) {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("!")
		return
	}
	if num < 0 {
		fmt.Println("F")
	} else if num > 0 {
		fmt.Println("P")
	} else if num == 0 {
		fmt.Println("0")
	} else {
		fmt.Printf("\n")
	}
	// return "\n"
}
