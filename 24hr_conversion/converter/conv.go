package converter

import (
	"strconv"
	"strings"
)

func TimeConversion(s string) string {
	// Write your code here
	// time := 0
	str := strings.Split(s, ":")

	num, err := strconv.Atoi(str[0])
	if err != nil {
		panic("INVALID STRING")
	}

	if strings.Contains(str[2], "PM") {
		if str[0] == "12" {
			str[2] = strings.Replace(str[2], "PM", "", 1)
			str[0] = strconv.Itoa(num)
		} else {
			str[2] = strings.Replace(str[2], "PM", "", 1)
			str[0] = strconv.Itoa(num + 12)
		}
	} else if strings.Contains(str[2], "AM") {
		if str[0] == "12" {
			str[2] = strings.Replace(str[2], "AM", "", 1)
			str[0] = strings.Replace(str[0], "12", "00", 1)
			// str[0] = strconv.Itoa(num - 12)
		} else {
			str[2] = strings.Replace(str[2], "AM", "", 1)
			str[0] = strconv.Itoa(num)
			str[0] = "0" + str[0]
		}
	}

	return str[0] + ":" + str[1] + ":" + str[2]
}
