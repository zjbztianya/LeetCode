package problem0098

import (
	"github.com/zjbztianya/LeetCode/kit"
	"math"
)

// TreeNode is pre-defined...
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
type TreeNode = kit.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func dfs(root *TreeNode) (bool, int, int) {
	if root == nil {
		return true, math.MaxInt64, math.MinInt64
	}
	lValid, lMin, lMax := dfs(root.Left)
	rValid, rMin, rMax := dfs(root.Right)
	if !lValid || !rValid {
		return false, 0, 0
	}
	if root.Val <= lMax || root.Val >= rMin {
		return false, 0, 0
	}
	return true, min(root.Val, lMin), max(root.Val, rMax)
}

func isValidBST(root *TreeNode) bool {
	ans, _, _ := dfs(root)
	return ans
}
