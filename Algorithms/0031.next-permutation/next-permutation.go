package problem0031

func Reverse(r []int) {
	n := len(r)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}
}

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i == -1 {
		Reverse(nums)
		return
	}
	j := i + 1
	for j < len(nums) && nums[i] < nums[j] {
		j++
	}
	nums[i], nums[j-1] = nums[j-1], nums[i]
	Reverse(nums[i+1:])
}
