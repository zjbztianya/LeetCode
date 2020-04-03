package problem0007

import "math"

func reverse(x int) int {
	sgn := 1
	y := int64(x)
	if x < 0 {
		sgn = -1
		y = -y
		if y > math.MaxInt32 {
			return 0
		}
	}

	var ans int64
	for y > 0 {
		ans = ans*10 + y%10
		y /= 10
	}
	ans *= int64(sgn)
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	}
	return int(ans)
}