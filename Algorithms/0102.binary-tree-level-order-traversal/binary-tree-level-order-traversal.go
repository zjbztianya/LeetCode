package problem0102

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

type Queque []interface{}

func (que *Queque) push(elem interface{}) {
	*que = append(*que, elem)
}

func (que *Queque) empty() bool {
	return len(*que) == 0
}
func (que *Queque) pop() interface{} {
	if que.empty() {
		return nil
	}
	front := (*que)[0]
	*que = (*que)[1:len(*que)]
	return front
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	que := new(Queque)
	que.push(root)
	que.push(nil)
	var level []int
	for !que.empty() {
		front, ok := que.pop().(*TreeNode)
		if ok {
			level = append(level, front.Val)
			if front.Left != nil {
				que.push(front.Left)
			}
			if front.Right != nil {
				que.push(front.Right)
			}
		} else {
			ans = append(ans, level)
			level = []int{}
			if !que.empty() {
				que.push(nil)
			}
		}
	}
	return ans
}
