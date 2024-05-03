package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func sumOfPrimes(num int) int {
	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func printRune(num int) {
	result := sumOfPrimes(num)
	fmt.Printf(" %d is: %d", num, result)
}

func main() {
	number := 1000000
	printRune(number)
}
