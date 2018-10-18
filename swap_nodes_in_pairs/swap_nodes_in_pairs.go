package swap_nodes_in_pairs

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{0, nil}
	p := dummyHead
	var n int
	var next *ListNode
	for ; head != nil; head = next {
		if (n & 1) == 0 {
			next = head.Next
			p.Next = head
			head.Next = nil
		} else {
			next = head.Next
			head.Next = p.Next
			p.Next = head
			p = head.Next
		}
		n++
	}
	return dummyHead.Next
}
