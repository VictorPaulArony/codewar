package acc

import "strings"

func Accum(words string) string {
	res := ""
	for i, word := range words {
		// change all the given chararcters into uppper case form first
		up := strings.ToUpper(string(word))
		// itterate through all the strings up to i
		rep := strings.Repeat(up, i)
		// change the itterated words to lower case
		low := strings.ToLower(rep)
		// concatination of the lower case and the upper case
		res += up + low
		// adding of the - to all except to the n-1 word
		if i != len(words)-1 {
			res += "-"
		}
	}
	return res
}
