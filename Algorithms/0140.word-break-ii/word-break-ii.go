package problem0140

func wordBreak(s string, wordDict []string) []string {
	n := len(s)
	if n == 0 {
		return []string{}
	}

	wordMap := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]] = true
	}

	dp := make([]bool, n+1)
	dp[n] = true
	path := make([][]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if _, ok := wordMap[s[i:j+1]]; ok && dp[j+1] {
				dp[i] = true
				path[i] = append(path[i], j+1)
			}
		}
	}
	var ans []string
	var dfs func(int, string)
	dfs = func(i int, str string) {
		if i == n {
			ans = append(ans, str[1:])
			return
		}
		for j := 0; j < len(path[i]); j++ {
			dfs(path[i][j], str+" "+s[i:path[i][j]])

		}
	}
	dfs(0, "")
	return ans
}
