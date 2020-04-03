package problem0042

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
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	left, right := make([]int, n), make([]int, n)
	left[0], right[n-1] = height[0], height[n-1]
	ans := 0
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i])
		right[n-i-1] = max(right[n-i], height[n-i-1])
	}

	for i := 0; i < n; i++ {
		t := min(left[i], right[i])
		if t > height[i] {
			ans += t - height[i]
		}
	}
	return ans
}
