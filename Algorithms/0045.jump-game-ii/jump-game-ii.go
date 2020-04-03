package problem0045

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
func jump(nums []int) int {
	var ans, pos int
	n := len(nums)
	for i := 0; i < n; i++ {
		if pos >= n-1 {
			break
		}
		j := pos
		for i <= j {
			pos = max(pos, i+nums[i])
			i++
		}
		i--
		ans++
	}
	return ans
}
