package intersection

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	var numA, numB int
//	cur := headA
//	for ; cur != nil; cur = cur.Next {
//		numA++
//	}
//	cur = headB
//	for ; cur != nil; cur = cur.Next {
//		numB++
//	}
//	var n int
//	if numA < numB {
//		n = numB - numA
//		for i := 0; i < n; i++ {
//			headB = headB.Next
//		}
//	} else {
//		n = numA - numB
//		for i := 0; i < n; i++ {
//			headA = headA.Next
//		}
//	}
//	for headA != nil && headB != nil && headA != headB {
//		headA = headA.Next
//		headB = headB.Next
//	}
//	if headA == headB {
//		return headA
//	}
//	return nil
//}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB
	for a != b && (a != nil || b != nil) {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
