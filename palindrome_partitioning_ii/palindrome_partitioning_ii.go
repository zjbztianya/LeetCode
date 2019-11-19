package palindrome_partitioning_ii

func minCut(s string) int {
	n := len(s)
	flag := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		flag[i] = make([]bool, n+1)
	}
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = n - i - 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 1 || flag[i+1][j-1]) {
				flag[i][j] = true
				if dp[i] > dp[j+1]+1 {
					dp[i] = dp[j+1] + 1
				}
			}
		}
	}
	return dp[0]
}
