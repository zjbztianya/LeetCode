package problem0018

import "sort"

func fourSum(nums []int, target int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			k, l := j+1, len(nums)-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})
					k, l = k+1, l-1
					for k < l && nums[k-1] == nums[k] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}
	}
	return ans
}

