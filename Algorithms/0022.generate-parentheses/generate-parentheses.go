package problem0022

func dfs(num, n int, s string, ans *[]string) {
	if len(s) == n {
		if num == 0 {
			*ans = append(*ans, s)
		}
		return
	}
	if num <= 0 {
		dfs(num-1, n, s+"(", ans)
	}
	if num < 0 {
		dfs(num+1, n, s+")", ans)
	}
}

func generateParenthesis(n int) []string {
	var ans []string
	dfs(0, 2*n, "", &ans)
	return ans
}
