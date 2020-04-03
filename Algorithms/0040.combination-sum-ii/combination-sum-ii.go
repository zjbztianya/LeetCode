package problem0040

import "sort"

func dfs(cur, target int, nums, vec []int, ans *[][]int) {
	if target == 0 {
		tmp := make([]int, len(vec))
		copy(tmp, vec)
		*ans = append(*ans, tmp)
		return
	}
	for i := cur; i < len(nums); i++ {
		if target < nums[i] || (i > cur && nums[i] == nums[i-1]) {
			continue
		}
		vec = append(vec, nums[i])
		dfs(i+1, target-nums[i], nums, vec, ans)
		vec = vec[:len(vec)-1]
	}
}
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
	var vec []int
	dfs(0, target, candidates, vec, &ans)
	return ans
}
