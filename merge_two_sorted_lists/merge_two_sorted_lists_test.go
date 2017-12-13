package merge_two_sorted_lists

import "testing"

func TestMergeTwoLists(t *testing.T) {
	num1, num2 := []int{1, 2, 4}, []int{1, 3, 4}
	l1 := new(ListNode)
	l2 := new(ListNode)

	p := l1
	p.Val = num1[0]
	for i := 1; i < len(num1); i++ {
		q := &ListNode{num1[i], nil}
		p.Next = q
		p = p.Next
	}

	p = l2
	p.Val = num2[0]
	for i := 1; i < len(num2); i++ {
		q := &ListNode{num2[i], nil}
		p.Next = q
		p = p.Next
	}

	ans := mergeTwoLists(l1, l2)
	output := []int{1, 1, 2, 3, 4, 4}
	for i := 0; i < len(output); i++ {
		if ans.Val != output[i] {
			t.Errorf("test case TestMergeTwoLists: output: %v want: %v", output, ans)
		}
		ans = ans.Next
	}
}
