package problem0059

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	i, j := 0, 0
	matrix[i][j] = 1
	for step := 1; step < n*n; {
		for j+1 < n && matrix[i][j+1] == 0 {
			step++
			matrix[i][j+1] = step
			j++
		}

		for i+1 < n && matrix[i+1][j] == 0 {
			step++
			matrix[i+1][j] = step
			i++
		}

		for j-1 >= 0 && matrix[i][j-1] == 0 {
			step++
			matrix[i][j-1] = step
			j--
		}

		for i-1 >= 0 && matrix[i-1][j] == 0 {
			step++
			matrix[i-1][j] = step
			i--
		}
	}
	return matrix
}
