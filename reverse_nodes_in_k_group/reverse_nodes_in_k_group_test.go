package reverse_nodes_in_k_group

import "testing"

func TestReverseKGroup1(t *testing.T) {
	var list *ListNode
	if got := reverseKGroup(list, 1); got != nil {
		t.Errorf("test case TestReverseKGroup1: output: %v want: %v", got, nil)
	}
}

func TestReverseKGroup2(t *testing.T) {
	list := &ListNode{1, nil}

	if got := reverseKGroup(list, 1); got != list {
		t.Errorf("test case TestReverseKGroup2: output: %v want: %v", got, list)
	}
}

func TestReverseKGroup3(t *testing.T) {
	num := []int{1, 2, 3, 4, 5}
	l := new(ListNode)

	p := l
	p.Val = num[0]
	for i := 1; i < len(num); i++ {
		q := &ListNode{num[i], nil}
		p.Next = q
		p = p.Next
	}

	ans := reverseKGroup(l, 3)
	want := []int{3, 2, 1, 4, 5}
	var i int
	for ans != nil {
		if ans.Val != want[i] {
			t.Errorf("test case TestReverseKGroup3 fail")
			//return
		}
		ans = ans.Next
		i++
	}
	if i != len(want) {
		t.Errorf("test case TestReverseKGroup3 fail")
	}
}
