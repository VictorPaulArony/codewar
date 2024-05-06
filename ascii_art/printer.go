func printer(str string) []string {
	var res []string
	if str == "" {
		fmt.Println("EMPTY STRING")
	}
	for _, char := range str {
		first := int(char-32)*9 - 2
		res = append(res, first)
	}
	fmt.Print(strings.Join(res, "\n"))

}