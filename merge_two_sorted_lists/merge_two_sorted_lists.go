package merge_two_sorted_lists

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{0, nil}
	p := dummyHead
	for l1 != nil || l2 != nil {
		if l1 == nil {
			p.Next = l2
			l2 = l2.Next
		} else if l2 == nil {
			p.Next = l1
			l1 = l1.Next
		} else if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	return dummyHead.Next
}
