package problem0074

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	l, r, row := 0, n-1, 0
	for l <= r {
		mid := (l + r) >> 1
		if matrix[mid][m-1] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l == n {
		return false
	}

	row, l, r = l, 0, m-1
	for l <= r {
		mid := (l + r) >> 1
		if matrix[row][mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l == m {
		return false
	}
	return matrix[row][l] == target
}
