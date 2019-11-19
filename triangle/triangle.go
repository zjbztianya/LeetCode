package triangle

import "math"

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	m := len(triangle[n-1])
	dp := make([][]int, 2)
	dp[0] = make([]int, m)
	dp[1] = make([]int, m)
	dp[0][0] = triangle[0][0]
	ans := math.MaxInt32
	for i := 1; i < n; i++ {
		dp[i&1][i] = dp[(i+1)&1][i-1] + triangle[i][i]
		dp[i&1][0] = dp[(i+1)&1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i&1][j] = min(dp[(i+1)&1][j-1], dp[(i+1)&1][j]) + triangle[i][j]
		}
	}
	for i := 0; i < m; i++ {
		ans = min(ans, dp[(n+1)&1][i])
	}
	return ans
}
