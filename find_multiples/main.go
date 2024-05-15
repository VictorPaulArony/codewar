package main

func FindMultiples(integer, limit int) []int {
	var res []int
	for i := integer; i <= limit; i++ {
		if i%integer == 0 {
			res = append(res, i)
		}
	}
	return res
}
