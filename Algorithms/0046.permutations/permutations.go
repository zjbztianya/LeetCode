package problem0046

func dfs(cur, n int, ans *[][]int, vec, nums []int, vis []bool) {
	if cur == n {
		tmp := make([]int, n)
		copy(tmp, vec)
		*ans = append(*ans, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if !vis[i] {
			vis[i] = true
			vec[cur] = nums[i]
			dfs(cur+1, n, ans, vec, nums, vis)
			vis[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	vec := make([]int, n)
	vis := make([]bool, n)
	dfs(0, n, &ans, vec, nums, vis)
	return ans
}
