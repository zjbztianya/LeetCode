package problem0089

func grayCode(n int) []int {
	ans := make([]int, 1<<uint(n))
	var m int
	for i := 0; i < n; i++ {
		m = 1 << uint(i)
		for j := m - 1; j >= 0; j-- {
			ans[2*m-j-1] = ans[j] | m
		}
	}
	return ans
}
