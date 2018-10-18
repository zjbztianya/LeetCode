//package subsets
package main

import "fmt"

func dfs(cur int, nums []int, ans *[][]int, a []int) {
	if cur >= len(nums) {
		tmp := make([]int, len(a))
		copy(tmp, a)
		*ans = append(*ans, tmp)
		return
	}
	dfs(cur+1, nums, ans, append(a, nums[cur]))
	dfs(cur+1, nums, ans, a)
}

func subsets(nums []int) [][]int {
	var ans [][]int
	dfs(0, nums, &ans, []int{})
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
