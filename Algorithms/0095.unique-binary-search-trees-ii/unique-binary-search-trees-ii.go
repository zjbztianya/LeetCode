package problem0095

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

func genBST(l int, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{nil}
	}
	var sub []*TreeNode
	for i := l; i <= r; i++ {
		lt := genBST(l, i-1)
		rt := genBST(i+1, r)
		for j := 0; j < len(lt); j++ {
			for k := 0; k < len(rt); k++ {
				node := new(TreeNode)
				node.Val, node.Left, node.Right = i, lt[j], rt[k]
				sub = append(sub, node)
			}
		}
	}
	return sub
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return genBST(1, n)
}
