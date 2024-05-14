package find

func FindOdd(seq []int) int {
	var res int
	count := make(map[int]int)
	for _, val := range seq {
		count[val]++
	}
	for i, v := range count {
		if v%2 != 0 {
			res = i
		}
	}
	return res
}
