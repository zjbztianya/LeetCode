package problem0139

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	if n == 0 {
		return true
	}

	wordMap := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = true
	}

	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			if _, ok := wordMap[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
