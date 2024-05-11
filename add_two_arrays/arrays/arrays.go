package arrays

func MergeArrays(arr1, arr2 []int) []int {
	var res []int
	// Determine the length of the result array
	length := len(arr1)
	if len(arr2) > length {
		length = len(arr2)
	}
	res = make([]int, length)
	// Add corresponding elements from both arrays
	for i := 0; i < length; i++ {
		if i < len(arr1) {
			// res[i] += arr1[i]
			res = append(res, arr1[i])
		}
		if i < len(arr2) {
			// res[i] += arr2[i]
			res = append(res, arr2[i])
		}
	}
	return res
}
