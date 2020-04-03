package problem0145

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
func dfs(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, ans)
	dfs(root.Right, ans)
	*ans = append(*ans, root.Val)
}

func postorderTraversal(root *TreeNode) (res []int) {
	dfs(root, &res)
	return res
}

