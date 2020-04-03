package problem0082


type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	myHead := new(ListNode)
	myHead.Next = head
	p := myHead
	var q *ListNode
	for p.Next != nil {
		if p.Next.Next != nil && p.Next.Val == p.Next.Next.Val {
			q = p.Next
			for q.Next != nil && q.Next.Val == p.Next.Val {
				q = q.Next
			}
			p.Next = q.Next
		} else {
			p = p.Next
		}
	}
	return myHead.Next
}
