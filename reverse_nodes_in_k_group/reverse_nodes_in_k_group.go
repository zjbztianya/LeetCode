package reverse_nodes_in_k_group

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var myHead, next *ListNode
	for ; head != nil; head = next {
		next = head.Next
		head.Next = myHead
		myHead = head
	}
	return myHead
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{0, head}
	p := dummyHead
	var next, tail *ListNode
	for i := 1; head != nil; i, head = i+1, next {
		next = head.Next
		if i%k == 0 {
			head.Next = nil
			q := reverseList(p.Next)
			p.Next.Next = next
			tail = p.Next
			p.Next = q
			p = tail
		}
	}
	return dummyHead.Next
}
