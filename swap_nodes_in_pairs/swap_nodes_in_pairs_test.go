package swap_nodes_in_pairs

import "testing"

func TestSwapPairs1(t *testing.T) {
	var list *ListNode
	if got := swapPairs(list); got != nil {
		t.Errorf("test case TestSwapPairs1: output: %v want: %v", got, nil)
	}
}

func TestSwapPairs2(t *testing.T) {
	list := &ListNode{1, nil}

	if got := swapPairs(list); got != list {
		t.Errorf("test case TestSwapPairs2: output: %v want: %v", got, list)
	}
}

func TestSwapPairs3(t *testing.T) {
	num := []int{1, 2, 3, 4}
	l := new(ListNode)

	p := l
	p.Val = num[0]
	for i := 1; i < len(num); i++ {
		q := &ListNode{num[i], nil}
		p.Next = q
		p = p.Next
	}

	ans := swapPairs(l)
	want := []int{2, 1, 4, 3}
	var i int
	for ans != nil {
		if ans.Val != want[i] {
			t.Errorf("test case TestSwapPairs3 fail")
		}
		ans = ans.Next
		i++
	}
	if i != len(want) {
		t.Errorf("test case TestSwapPairs3 fail")
	}
}
