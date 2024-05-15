package sample

import "strings"

func Initial(name string) string {
	// Split the name into first and last name
	parts := strings.Split(name, " ")
	// Extract the first letter of each part and convert to uppercase
	initials := strings.ToUpper(string(parts[0][0])) + "." + strings.ToUpper(string(parts[1][0]))
	return initials
}
