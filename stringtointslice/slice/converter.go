package slice

func StringToIntSlice(str string) []int {
	var res []int
	for _, s := range str {
		res = append(res, int(s))
	}
	return res
}
