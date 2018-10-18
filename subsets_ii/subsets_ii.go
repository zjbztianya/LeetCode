package main

import "sort"

func dfs(cur int, nums []int, ans *[][]int, a []int) {
	if cur >= len(nums) {
		return
	}
	for i := cur; i < len(nums); i++ {
		if i-1 >= cur && nums[i-1] == nums[i] {
			continue
		}
		a = append(a, nums[i])
		dfs(i+1, nums, ans, a)
		tmp := make([]int, len(a))
		copy(tmp, a)
		*ans = append(*ans, tmp)
		a = a[:len(a)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	ans = append(ans, []int{})
	dfs(0, nums, &ans, []int{})
	return ans
}
