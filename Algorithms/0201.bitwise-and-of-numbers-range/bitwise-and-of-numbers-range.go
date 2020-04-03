package problem0201

func rangeBitwiseAnd(m int, n int) int {
	for ; n > m; n &= n - 1 {
	}
	return n
}
