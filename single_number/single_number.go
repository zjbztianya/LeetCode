package single_number

func singleNumber(nums []int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}
