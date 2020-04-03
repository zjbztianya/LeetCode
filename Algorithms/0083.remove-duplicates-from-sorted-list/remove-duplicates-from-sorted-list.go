package problem0083

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	myHead := new(ListNode)
	myHead.Next = head
	p := myHead
	for ; p.Next != nil; p = p.Next {
		q := p.Next
		for q.Next != nil && q.Val == q.Next.Val {
			q = q.Next
		}
		p.Next = q
	}
	return myHead.Next
}

