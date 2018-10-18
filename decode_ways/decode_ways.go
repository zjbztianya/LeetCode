package decode_ways

func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	if s[0] == '0' {
		return 0
	}
	for i := 1; i < n; i++ {
		if s[i] > '0' {
			dp[i+1] = dp[i]
		}
		num := (s[i-1]-'0')*10 + s[i] - '0'
		if num >= 10 && num <= 26 {
			dp[i+1] += dp[i-1]
		}
	}
	return dp[n]
}
