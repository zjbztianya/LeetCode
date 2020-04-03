package problem0130

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func dfs(x int, y int, board [][]byte, n int, m int) {
	board[x][y] = '$'
	for i := 0; i < 4; i++ {
		xx, yy := x+dx[i], y+dy[i]
		if xx < 0 || xx >= n || yy < 0 || yy >= m {
			continue
		}
		if board[xx][yy] == 'O' {
			dfs(xx, yy, board, n, m)
		}
	}
}

func solve(board [][]byte) {
	n := len(board)
	if n == 0 {
		return
	}
	m := len(board[0])
	for i := 0; i < m; i++ {
		if board[0][i] == 'O' {
			dfs(0, i, board, n, m)
		}
		if board[n-1][i] == 'O' {
			dfs(n-1, i, board, n, m)
		}
	}

	for i := 1; i < n-1; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0, board, n, m)
		}
		if board[i][m-1] == 'O' {
			dfs(i, m-1, board, n, m)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '$' {
				board[i][j] = 'O'
			}
		}
	}
}
