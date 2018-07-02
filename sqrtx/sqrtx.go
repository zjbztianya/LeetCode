package sqrtx

func mySqrt(x int) int {
	l, r := 0, x
	var mid int
	for l <= r {
		mid = (l + r) >> 1
		if mid*mid <= x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}
