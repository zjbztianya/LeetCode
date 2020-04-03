package problem0198

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n+5)
	dp[1] = nums[0]
	for i := 1; i < n; i++ {
		dp[i+1] = max(dp[i-1]+nums[i], dp[i])
	}
	return dp[n]
}
