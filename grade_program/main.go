package main

import "grade_program/grades"

func main() {
	// Test cases
	grade := grades.GetGrade(85, 90, 88)
	println(grade) // Output: B
}
