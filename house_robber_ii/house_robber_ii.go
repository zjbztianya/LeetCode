package house_robber_ii

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
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n+5)
	dp[1] = nums[0]
	for i := 1; i < n-1; i++ {
		dp[i+1] = max(dp[i-1]+nums[i], dp[i])
	}
	ans := dp[n-1]
	dp = make([]int, n+5)
	dp[2] = nums[1]
	for i := 2; i < n; i++ {
		dp[i+1] = max(dp[i-1]+nums[i], dp[i])
	}
	ans = max(ans, dp[n])
	return ans
}
