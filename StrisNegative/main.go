package main

import "StrisNegative/negative"

func main() {
	negative.StrIsNegative("585")
	negative.StrIsNegative("-58")
	negative.StrIsNegative("55s44")
	negative.StrIsNegative("101-1331")
	negative.StrIsNegative("5544-")
}
