package problem0080

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	pos := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[pos-2] {
			continue
		}
		nums[pos] = nums[i]
		pos++
	}
	return pos
}
