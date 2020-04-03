package problem0147

import "github.com/zjbztianya/LeetCode/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = kit.ListNode

func insertionSortList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	pre, cur, next := dummy, head, (*ListNode)(nil)
	for cur != nil {
		next = cur.Next
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur

		cur = next
		pre = dummy
	}
	return dummy.Next
}
