package problem0106

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

func makeTree(postorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	val, n := postorder[len(postorder)-1], 0
	for inorder[n] != val {
		n++
	}
	node := &TreeNode{val, nil, nil}
	node.Left = makeTree(postorder[0:n], inorder[0:n])
	node.Right = makeTree(postorder[n:len(postorder)-1], inorder[n+1:])
	return node
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return makeTree(postorder, inorder)
}
