package problem0047

import "sort"

func dfs(cur, n int, ans *[][]int, vec, nums []int, vis []bool) {
	if cur == n {
		tmp := make([]int, n)
		copy(tmp, vec)
		*ans = append(*ans, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if !vis[i] {
			if i-1 >= 0 && !vis[i-1] && nums[i-1] == nums[i] {
				continue
			}
			vis[i] = true
			vec[cur] = nums[i]
			dfs(cur+1, n, ans, vec, nums, vis)
			vis[i] = false
		}
	}
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int
	vec := make([]int, n)
	vis := make([]bool, n)
	dfs(0, n, &ans, vec, nums, vis)
	return ans
}
