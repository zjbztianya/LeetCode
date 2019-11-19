package trailing_zeroes
//172. Factorial Trailing Zeroes

func trailingZeroes(n int) int {
	var n5 int
	for n > 0 {
		n /= 5
		n5 += n
	}
	return n5
}
