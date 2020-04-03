package problem0055

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
func canJump(nums []int) bool {
	var pos int
	n := len(nums)
	for i := 0; i < n; i++ {
		pos = max(pos, i+nums[i])
		if pos <= i && i != n-1 {
			return false
		}
	}
	return pos >= n-1
}
