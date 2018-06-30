package unique_paths_ii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		if obstacleGrid[i][0] == 1 {
			continue
		}
		if i > 0 {
			dp[i][0] = dp[i-1][0]
		} else {
			dp[i][0] = 1
		}
	}

	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] & (obstacleGrid[0][i] ^ 1)
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}
