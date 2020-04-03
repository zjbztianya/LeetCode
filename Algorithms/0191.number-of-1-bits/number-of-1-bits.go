package problem0191

func hammingWeight(num uint32) int {
	var n int
	for num > 0 {
		n++
		num &= num - 1
	}
	return n
}
