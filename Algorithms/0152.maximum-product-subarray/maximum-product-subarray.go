package problem0152

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	maxN, minN, ans := nums[0], nums[0], nums[0]
	var tmp int
	for i := 1; i < len(nums); i++ {
		tmp = maxN
		maxN = max(minN*nums[i], max(maxN*nums[i], nums[i]))
		minN = min(tmp*nums[i], min(minN*nums[i], nums[i]))
		ans = max(ans, maxN)
	}
	return ans
}
