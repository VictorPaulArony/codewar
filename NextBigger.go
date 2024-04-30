func NextBigger(num int) int {
    // Convert the number to a slice of digits
    digits := make([]int, 0)
    for num > 0 {
        digits = append([]int{num % 10}, digits...)
        num /= 10
    }

    // Find the pivot point
    pivot := -1
    for i := len(digits) - 2; i >= 0; i-- {
        if digits[i] < digits[i+1] {
            pivot = i
            break
        }
    }

    if pivot == -1 {
        // No pivot point found, return -1
        return -1
    }

    // Find the smallest digit to the right of the pivot that is greater than the pivot
    next := pivot + 1
    for i := pivot + 2; i < len(digits); i++ {
        if digits[i] > digits[pivot] && digits[i] < digits[next] {
            next = i
        }
    }

    // Swap the pivot and the next digit
    digits[pivot], digits[next] = digits[next], digits[pivot]

    // Reverse the digits to the right of the pivot
    for i, j := pivot+1, len(digits)-1; i < j; i, j = i+1, j-1 {
        digits[i], digits[j] = digits[j], digits[i]
    }

    // Convert the slice of digits back to an integer
    result := 0
    for _, digit := range digits {
        result = result*10 + digit
    }

    return result
}
