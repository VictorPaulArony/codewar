package rev

func StrRev(str string) string {
	//     res := ""
	// for i := len(str)-1; i >=  0; i-- {
	//   res += string(str[i])
	// }
	// return
	res := []rune(str)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
