package range_bitwise_and

func rangeBitwiseAnd(m int, n int) int {
	for ; n > m; n &= n - 1 {
	}
	return n
}
