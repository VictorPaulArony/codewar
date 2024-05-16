package mid

func Middle(s string) string {
	length := len(s)

	// Check if the word's length is odd or even
	if length%2 == 0 { // Even length
		middle := length / 2
		return s[middle-1 : middle+1]
	} else { // Odd length
		middle := length / 2
		return string(s[middle])
	}
}
