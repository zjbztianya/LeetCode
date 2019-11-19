package palindrome_partitioning

var ans [][]string
var a []string

func ok(i int, j int, s string) bool {
	for ; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func dfs(pos int, n int, s string) {
	if pos == n {
		tmp := make([]string, len(a))
		copy(tmp, a)
		ans = append(ans, tmp)
		return
	}
	for i := pos; i < n; i++ {
		if ok(pos, i, s) {
			a = append(a, s[pos:i+1])
			dfs(i+1, n, s)
			a = a[:len(a)-1]
		}
	}
}

func partition(s string) [][]string {
	a = a[:0]
	ans = ans[:0]
	dfs(0, len(s), s)
	return ans
}
