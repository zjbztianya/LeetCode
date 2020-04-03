package problem0092

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	myHead := new(ListNode)
	myHead.Next = head
	prev := myHead

	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}

	cur := prev.Next
	var nxt *ListNode
	for i := m - 1; i < n-1; i++ {
		nxt = cur.Next
		cur.Next = cur.Next.Next
		nxt.Next = prev.Next
		prev.Next = nxt
	}
	return myHead.Next
}
