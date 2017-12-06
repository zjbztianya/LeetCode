package string_to_integer_atoi

import "math"

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	var num int64
	var i int
	sgn := 1
	for str[i] == ' ' {
		i++
	}
	if str[i] == '-' {
		sgn = -1
		i++
	} else if str[i] == '+' {
		i++
	}
	str = str[i:]

	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			break
		}
		num = num*10 + (int64(str[i] - '0'))
		if sgn > 0 && num > math.MaxInt32 {
			return math.MaxInt32
		}
		if sgn < 0 && num > -math.MinInt32 {
			return math.MinInt32
		}
	}
	return int(num) * sgn
}