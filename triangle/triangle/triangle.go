package triangle

func Triangle(a, b, c int) bool {
	// Check if all sides are greater than zero
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	// Check the triangle inequality theorem
	return (a+b > c) && (a+c > b) && (b+c > a)
}

// func IsTriangle(a, b, c int) bool {
//   return a+b > c && b+c > a && a+c > b
// }
