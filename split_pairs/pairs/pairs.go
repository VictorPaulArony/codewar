package pairs

import "math"

func Solution(s string) (r []string) {
	var res []string
	l := len(s)
	if math.Mod(float64(l), 2) != 0 {
		s += "_"
	}
	for i := 0; i < l; i += 2 {
		res = append(res, s[i:i+2])
	}
	return res
}
