package main

func GetSize(w, h, d int) [2]int {
	var area int
	var vol int
	var res [2]int
	area = 2*w*h + 2*h*d + 2*w*d
	vol = w * h * d
	res = [2]int{area, vol}
	return res
}
