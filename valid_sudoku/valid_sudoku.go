package valid_sudoku

const n = 9

func isValidSudoku(board [][]byte) bool {
	var row [n][n]bool
	var col [n][n]bool
	var region [n][n]bool
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '0' - 1
			key := i/3*3 + j/3
			if row[i][num] || col[j][num] || region[key][num] {
				return false
			}
			row[i][num], col[j][num], region[key][num] = true, true, true
		}
	}
	return true
}
