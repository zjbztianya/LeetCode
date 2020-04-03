package problem0200

var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, -1, 1}

func dfs(x, y, n, m int, grid [][]byte) {
	grid[x][y] = '0'
	for i := 0; i < 4; i++ {
		xx, yy := x+dx[i], y+dy[i]
		if xx < 0 || xx >= n || yy < 0 || yy >= m {
			continue
		}
		if grid[xx][yy] == '1' {
			dfs(xx, yy, n, m, grid)
		}
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var ans int
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				dfs(i, j, n, m, grid)
				ans++
			}
		}
	}
	return ans
}
