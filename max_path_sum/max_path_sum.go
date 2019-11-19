package max_path_sum

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftSum := max(0, dfs(root.Left))
	rightSum := max(0, dfs(root.Right))
	ans = max(ans, leftSum+rightSum+root.Val)
	return max(leftSum, rightSum) + root.Val
}

func maxPathSum(root *TreeNode) int {
	ans = math.MinInt32
	dfs(root)
	return ans
}
