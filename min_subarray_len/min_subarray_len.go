package min_subarray_len

import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minSubArrayLen(s int, nums []int) int {
	ans, sum, left := math.MaxInt64, 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= s {
			ans = min(ans, i-left+1)
			sum -= nums[left]
			left++
		}
	}

	if ans == math.MaxInt64 {
		return 0
	} else {
		return ans
	}
}
