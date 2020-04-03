package problem0189

func reverse(r []int) {
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}
}

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	}
	k = k % len(nums)
	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:])
}
