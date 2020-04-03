package problem0109

import "github.com/zjbztianya/LeetCode/kit"

// TreeNode is pre-defined...
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
type TreeNode = kit.TreeNode
type ListNode = kit.ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


var myhead *ListNode

func buildTree(l int, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	node := new(TreeNode)
	node.Left = buildTree(l, mid-1)
	node.Val = myhead.Val
	myhead = myhead.Next
	node.Right = buildTree(mid+1, r)
	return node

}

func sortedListToBST(head *ListNode) *TreeNode {
	var n int
	for p := head; p != nil; p = p.Next {
		n++
	}
	myhead = head
	return buildTree(0, n-1)
}