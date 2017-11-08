package remove_duplicates_from_sorted_array

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
