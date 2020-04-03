package problem0026

func removeDuplicates(nums []int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		nums[n] = nums[i]
		n++
	}
	return n
}
