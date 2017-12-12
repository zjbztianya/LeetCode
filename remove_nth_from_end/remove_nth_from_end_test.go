package remove_nth_from_end

import "testing"

func TestRemoveNthFromEndFirst(t *testing.T) {
	num, n := []int{1, 2, 3, 4, 5}, 2
	p := new(ListNode)
	head := p
	p.Val = num[0]
	for i := 1; i < len(num); i++ {
		q := &ListNode{num[i], nil}
		p.Next = q
		p = p.Next
	}
	ans := removeNthFromEnd(head, n)
	output := []int{1, 2, 3, 5}
	for i := 0; i < len(output); i++ {
		if ans.Val != output[i] {
			t.Errorf("test case RemoveNthFromEndFirst: output: %v ans: %v", output, ans)
		}
		ans = ans.Next
	}
}

func TestRemoveNthFromEndSecond(t *testing.T) {
	num, n := []int{1}, 1
	p := new(ListNode)
	head := p
	p.Val = num[0]
	for i := 1; i < len(num); i++ {
		q := &ListNode{num[i], nil}
		p.Next = q
		p = p.Next
	}
	ans := removeNthFromEnd(head, n)
	if ans != nil {
		t.Errorf("test case TestRemoveNthFromEndSecond: output shuld nil  but ans not nil")
	}
}
