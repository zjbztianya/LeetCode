package problem0202

func isHappy(n int) bool {
	for n != 1 && n != 4 {
		var sum int
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
	}
	return n == 1
}
