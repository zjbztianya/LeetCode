package problem0216

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var dfs func(cur int, sum int, vec []int)
	dfs = func(cur int, sum int, vec []int) {
		if len(vec) == k {
			if sum == n {
				tmp := make([]int, k)
				copy(tmp, vec)
				ans = append(ans, tmp)
			}
			return
		}
		if cur == 10 {
			return
		}
		dfs(cur+1, sum, vec)
		vec = append(vec, cur)
		dfs(cur+1, sum+cur, vec)
		vec = vec[:len(vec)-1]
	}
	dfs(1, 0, []int{})
	return ans
}
