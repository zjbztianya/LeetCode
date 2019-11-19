package sort_list

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getMid(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeList(left, right *ListNode) *ListNode {
	dummy := new(ListNode)
	cur := dummy
	for left != nil && right != nil {
		if left.Val <= right.Val {
			cur.Next, left = left, left.Next
		} else {
			cur.Next, right = right, right.Next
		}
		cur = cur.Next
	}
	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMid(head)
	next := mid.Next
	mid.Next = nil
	left, right := sortList(head), sortList(next)
	return mergeList(left, right)
}
