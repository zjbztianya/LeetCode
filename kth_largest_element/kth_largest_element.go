package kth_largest_element

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	mid := nums[n>>1]
	l, r := 0, n-1
	for l <= r {
		for nums[l] > mid {
			l++
		}
		for nums[r] < mid {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	if r+1 >= k {
		return findKthLargest(nums[:r+1], k)
	} else {
		return findKthLargest(nums[r+1:], k-r-1)
	}
}
