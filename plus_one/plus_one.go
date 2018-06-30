package plus_one

func plusOne(digits []int) []int {
	digits = append(digits, 1)
	for i := len(digits) - 1; i > 0; i-- {
		digits[i] = digits[i] + digits[i-1]
		digits[i-1] = digits[i] / 10
		digits[i] %= 10
	}
	if digits[0] == 0 {
		return digits[1:]
	} else {
		return digits
	}
}
