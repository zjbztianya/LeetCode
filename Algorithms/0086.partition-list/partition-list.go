package problem0086

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h1, h2 := new(ListNode), new(ListNode)
	p1, p2, p := h1, h2, head
	for p != nil {
		tmp := p.Next
		if p.Val < x {
			p1.Next, p1 = p, p
		} else {
			p2.Next, p2 = p, p
		}
		p = tmp
	}
	p1.Next = h2.Next
	p2.Next = nil
	return h1.Next
}

