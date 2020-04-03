package problem0110

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