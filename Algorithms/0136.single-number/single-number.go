package problem0136

func singleNumber(nums []int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}

