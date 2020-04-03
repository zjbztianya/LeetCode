package problem0029

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}
	sgn := -1
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		sgn = 1
	}
	a, b := int64(dividend), int64(divisor)
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	ans := 0
	for a >= b {
		for i, num := 1, b; num <= a; i, num = i<<1, num<<1 {
			ans += i
			a -= num
		}
	}
	return ans * sgn
}
