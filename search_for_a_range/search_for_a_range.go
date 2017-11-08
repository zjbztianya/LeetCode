package search_for_a_range

func Binary(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	l := Binary(nums, target)
	if l != len(nums) && nums[l] == target {
		ans[0] = l
	}

	if ans[0] != -1 {
		ans[1] = Binary(nums, target+1) - 1
	}
	return ans
}
