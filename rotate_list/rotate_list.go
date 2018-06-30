package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	p := head
	n := 1
	for p.Next != nil {
		n++
		p = p.Next
	}
	k %= n
	if k == 0 {
		return head
	}
	tail := p

	q := head
	for i := 0; i < k; i++ {
		q = q.Next
	}
	p = head
	for q.Next != nil {
		q = q.Next
		p = p.Next
	}
	tail.Next = head
	head = p.Next
	p.Next = nil
	return head
}
