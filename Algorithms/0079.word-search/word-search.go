package problem0079


var dx = []int{0, 0, -1, 1}
var dy = []int{-1, 1, 0, 0}

func dfs(x int, y int, step int, board [][]byte, word string) bool {
	if step >= len(word) {
		return true
	}

	tmp := board[x][y]
	board[x][y] = '0'
	n, m := len(board), len(board[0])
	for d := 0; d < 4; d++ {
		tx, ty := x+dx[d], y+dy[d]
		if tx < 0 || tx >= n || ty < 0 || ty >= m {
			continue
		}

		if board[tx][ty] == word[step] {
			if dfs(tx, ty, step+1, board, word) {
				return true
			}

		}
	}
	board[x][y] = tmp
	return false
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	n, m := len(board), len(board[0])
	if n*m < len(word) {
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] && dfs(i, j, 1, board, word) {
				return true
			}
		}
	}
	return false
}