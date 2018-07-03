package combinations

func dfs(cur int, n int, k int, ans *[][]int, a []int) {
	if len(a) == k {
		tmp := make([]int, len(a))
		copy(tmp, a)
		*ans = append(*ans, tmp)
		return
	}
	if cur > n {
		return
	}
	dfs(cur+1, n, k, ans, append(a, cur))
	dfs(cur+1, n, k, ans, a)
}

func combine(n int, k int) [][]int {
	var ans [][]int
	dfs(1, n, k, &ans, []int{})
	return ans
}
