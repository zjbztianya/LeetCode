package problem0112

import "github.com/zjbztianya/LeetCode/kit"

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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return (sum - root.Val) == 0
	}

	l := hasPathSum(root.Left, sum-root.Val)
	r := hasPathSum(root.Right, sum-root.Val)
	return l || r
}