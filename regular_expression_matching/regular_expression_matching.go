package regular_expression_matching

//func isMatch(s string, p string) bool {
//	n, m := len(s), len(p)
//	if m == 0 {
//		return n == 0
//	}
//	flag := n > 0 && (s[0] == p[0] || p[0] == '.')
//	if m >= 2 && p[1] == '*' {
//		return (flag && isMatch(s[1:], p)) || isMatch(s, p[2:])
//	} else {
//		return flag && isMatch(s[1:], p[1:])
//	}
//}
type Mem struct {
	dp [][]bool
}

func (mem *Mem) dfs(i int, j int, s string, p string) bool {
	if j == len(p) {
		mem.dp[i][j] = i == len(s)
		return mem.dp[i][j]
	}
	flag := i < len(s) && (s[i] == p[j] || p[j] == '.')
	if j+2 <= len(p) && p[j+1] == '*' {
		mem.dp[i][j] = mem.dfs(i, j+2, s, p) || (flag && mem.dfs(i+1, j, s, p))
	} else {
		mem.dp[i][j] = flag && mem.dfs(i+1, j+1, s, p)
	}
	return mem.dp[i][j]
}

func isMatch(s string, p string) bool {
	mem := Mem{make([][]bool, len(s)+1)}
	for i := range mem.dp {
		mem.dp[i] = make([]bool, len(p)+1)
	}
	return mem.dfs(0, 0, s, p)
}
