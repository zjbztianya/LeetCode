package maximum_gap

import "math"

func memset(a []int, v int) {
	for i := range a {
		a[i] = v
	}
}

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	mx, mi := nums[0], nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > mx {
			mx = nums[i]
		}
		if nums[i] < mi {
			mi = nums[i]
		}
	}

	bucketCap := (mx - mi + n - 2) / (n - 1)
	if bucketCap == 0{
		return 0
	}
	bucketNum := (mx-mi)/bucketCap + 1
	bucketMax := make([]int, bucketNum)
	memset(bucketMax, math.MinInt64)
	bucketMin := make([]int, bucketNum)
	memset(bucketMin, math.MaxInt64)
	for i := 0; i < n; i++ {
		index := (nums[i] - mi) / bucketCap
		if bucketMin[index] > nums[i] {
			bucketMin[index] = nums[i]
		}
		if bucketMax[index] < nums[i] {
			bucketMax[index] = nums[i]
		}
	}
	res := math.MinInt64
	pre := 0
	for i := 1; i < bucketNum; i++ {
		if bucketMax[pre] == math.MinInt64 || bucketMin[i] == math.MaxInt64 {
			continue
		}
		diff := bucketMin[i] - bucketMax[pre]
		if diff > res {
			res = diff
		}
		pre = i
	}
	return res
}
