package largest_rectangle_area

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	seqL := make([]int, n)
	seqR := make([]int, n)
	stack := make([]int, n+1)
	top := 0
	for i := 0; i < n; i++ {
		for top > 0 && heights[stack[top]] >= heights[i] {
			top--
		}
		if top > 0 {
			seqL[i] = stack[top]
		} else {
			seqL[i] = -1
		}
		top++
		stack[top] = i
	}
	top = 0
	for i := n - 1; i >= 0; i-- {
		for top > 0 && heights[stack[top]] >= heights[i] {
			top--
		}
		if top > 0 {
			seqR[i] = stack[top]
		} else {
			seqR[i] = n
		}
		top++
		stack[top] = i
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, heights[i]*(seqR[i]-seqL[i]-1))
	}
	return ans
}
