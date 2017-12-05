package maximum_average_subarray

import "math"

func findMaxAverage(nums []int, k int) float64 {
	ans, sum := float64(math.MinInt64), 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i >= k-1 {
			ans = math.Max(ans, float64(sum)/float64(k))
			sum -= nums[i-k+1]
		}
	}
	return ans
}
