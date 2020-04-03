package problem0153

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := int(uint(l+r) >> 1)
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
