package problem0027

func removeElement(nums []int, val int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[n] = nums[i]
			n++
		}
	}
	nums = nums[:n]
	return n
}
