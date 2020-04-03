package problem0143

import "github.com/zjbztianya/LeetCode/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = kit.ListNode

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	cur := slow.Next
	slow.Next = nil
	mid := (*ListNode)(nil)
	for cur != nil {
		next := cur.Next
		cur.Next = mid
		mid = cur
		cur = next
	}

	dummy := new(ListNode)
	p, q := dummy, mid
	for head != nil || q != nil {
		p.Next = head
		head = head.Next
		p = p.Next
		if q != nil {
			p.Next = q
			q = q.Next
			p = p.Next
		}
	}
	head = dummy.Next
}

