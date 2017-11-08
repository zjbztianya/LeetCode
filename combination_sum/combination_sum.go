package combination_sum

func dfs(cur, target int, nums, vec []int, ans *[][]int) {
	if target == 0 {
		tmp := make([]int, len(vec))
		copy(tmp, vec)
		*ans = append(*ans, tmp)
		return
	}

	if cur == len(nums) {
		return
	}
	dfs(cur+1, target, nums, vec, ans)
	if target >= nums[cur] {
		vec = append(vec, nums[cur])
		dfs(cur, target-nums[cur], nums, vec, ans)
	}
}
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var vec []int
	dfs(0, target, candidates, vec, &ans)
	return ans
}
