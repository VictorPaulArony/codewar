package min_max

func MinMax(arr []int) (int, int) {
	min := arr[0]
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}
	return min, max
}
