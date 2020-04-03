package problem0050

func myPow(x float64, n int) float64 {
	sgn := 1
	if n < 0 {
		sgn = -1
		n = -n
	}
	ans, a := 1.0, x
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ans *= a
		}
		a *= a
	}
	if sgn == -1 {
		return 1.0 / ans
	}
	return ans
}