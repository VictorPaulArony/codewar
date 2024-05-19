package vowels

func CountVowels(str string) (count int) {
	vol := string("aeiou")
	for _, s := range str {
		for _, char := range vol {
			if char == s {
				count++
			}
		}
	}
	return count
}
