package balanced_tree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func dfs(root *TreeNode, ans *bool) int {
	if root == nil {
		return 0
	}
	dl, dr := dfs(root.Left, ans), dfs(root.Right, ans)
	if math.Abs(float64(dl-dr)) > 1 {
		*ans = false
	}
	return max(dl, dr) + 1
}
func isBalanced(root *TreeNode) bool {
	ans := true
	dfs(root, &ans)
	return ans
}
