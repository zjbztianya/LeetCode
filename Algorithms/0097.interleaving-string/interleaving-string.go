package problem0097

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m := len(s1), len(s2)
	if n+m != len(s3) {
		return false
	}
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		dp[0][i] = dp[0][i-1] && (s2[i-1] == s3[i-1])
	}

	for i := 1; i <= n; i++ {
		dp[i][0] = dp[i-1][0] && (s1[i-1] == s3[i-1])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i+1][j+1] = dp[i][j+1] == true && s1[i] == s3[i+j+1]
			dp[i+1][j+1] = dp[i+1][j+1] || (dp[i+1][j] == true && s2[j] == s3[i+j+1])
		}
	}
	return dp[n][m]
}
