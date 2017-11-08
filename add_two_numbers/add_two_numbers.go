package add_two_numbers

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	m := 0
	head := new(ListNode)
	tail := head
	for l1 != nil || l2 != nil {
		p := new(ListNode)
		if l1 != nil {
			p.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			p.Val += l2.Val
			l2 = l2.Next
		}

		p.Val += m
		m = p.Val / 10
		p.Val = p.Val % 10
		tail.Next = p
		tail = p
	}
	if m > 0 {
		tail.Next = &ListNode{m, nil}
	}
	return head.Next
}
