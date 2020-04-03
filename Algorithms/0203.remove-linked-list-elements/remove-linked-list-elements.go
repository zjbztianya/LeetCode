package problem0203

import "github.com/zjbztianya/LeetCode/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = kit.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{0, head}
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
