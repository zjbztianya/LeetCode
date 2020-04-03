package problem0054

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	n := len(matrix[0])
	m := len(matrix)
	var ans []int
	i, j := 0, 0
	ans = append(ans, matrix[i][j])
	flag := make([]bool, n*m)
	flag[0] = true
	for step := 1; step < n*m; {
		for j+1 < n && !flag[i*n+j+1] {
			ans = append(ans, matrix[i][j+1])
			flag[i*n+j+1] = true
			step++
			j++
		}

		for i+1 < m && !flag[(i+1)*n+j] {
			ans = append(ans, matrix[i+1][j])
			flag[(i+1)*n+j] = true
			step++
			i++
		}

		for j-1 >= 0 && !flag[(i*n+j-1)] {
			ans = append(ans, matrix[i][j-1])
			flag[(i*n + j - 1)] = true
			step++
			j--
		}

		for i-1 >= 0 && !flag[((i-1)*n+j)] {
			ans = append(ans, matrix[i-1][j])
			flag[((i-1)*n + j)] = true
			step++
			i--
		}
	}
	return ans
}
