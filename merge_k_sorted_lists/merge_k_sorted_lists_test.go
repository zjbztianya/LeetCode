package merge_k_sorted_lists

import (
	"reflect"
	"testing"
)

func TestMergeKLists1(t *testing.T) {
	var lists []*ListNode
	if got := mergeKLists(lists); got != nil {
		t.Errorf("test case TestMergeKLists1: output: %v want: %v", got, nil)
	}
}

func TestMergeKLists2(t *testing.T) {
	lists := []*ListNode{nil}

	if got := mergeKLists(lists); got != nil {
		t.Errorf("test case TestMergeKLists2: output: %v want: %v", got, nil)
	}
}

func TestMergeKList3(t *testing.T) {
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
	lists := []*ListNode{l1, l2}
	ans := mergeKLists(lists)
	want := []int{1, 1, 2, 3, 4, 4}
	var num []int
	for ans != nil {
		num = append(num, ans.Val)
		ans = ans.Next
	}
	if !reflect.DeepEqual(num, want) {
		t.Errorf("test case TestMergeTwoLists: output: %v want: %v", num, want)
	}
}
