package maximal_rectangle

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	n, m := len(matrix), len(matrix[0])
	h := make([]int, m+1)
	h[0] = int(matrix[0][0] - '0')
	cnt, ans := h[0], h[0]
	for i := 1; i < m; i++ {
		if matrix[0][i] == '0' {
			h[i], cnt = 0, 0
		} else {
			h[i], cnt = 1, cnt+1
			ans = max(ans, cnt)
		}
	}
	stack := make([]int, m+1)
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				h[j] = 0
			} else {
				h[j] += 1
			}
		}
		top := -1
		for j := 0; j < m+1; {
			if top == -1 || h[stack[top]] < h[j] {
				top++
				stack[top] = j
				j++
			} else {
				t := stack[top]
				top--
				if top == -1 {
					ans = max(ans, h[t]*j)
				} else {
					ans = max(ans, h[t]*(j-stack[top]-1))
				}
			}
		}
	}
	return ans
}
