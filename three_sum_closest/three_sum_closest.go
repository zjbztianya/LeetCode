package three_sum_closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-ans)) {
				ans = sum
			}
			if sum == target {
				return ans
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}
