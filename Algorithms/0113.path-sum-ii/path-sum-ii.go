package problem0113

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
func dfs(root *TreeNode, sum int, ans *[][]int, value []int) {
	if root == nil {
		return
	}

	value = append(value, root.Val)
	if root.Left == nil && root.Right == nil && (sum-root.Val) == 0 {
		tmp := make([]int, len(value))
		copy(tmp, value)
		*ans = append(*ans, tmp)
		return
	}

	dfs(root.Left, sum-root.Val, ans, value)
	dfs(root.Right, sum-root.Val, ans, value)
	value = value[:len(value)-1]
}

func pathSum(root *TreeNode, sum int) [][]int {
	var ans [][]int
	var value []int
	dfs(root, sum, &ans, value)
	return ans
}