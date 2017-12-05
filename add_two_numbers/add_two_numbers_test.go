package add_two_numbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 2
	l12 := new(ListNode)
	l12.Val = 4
	l1.Next = l12
	l13 := new(ListNode)
	l13.Val = 3
	l12.Next = l13

	l2 := new(ListNode)
	l2.Val = 5
	l22 := new(ListNode)
	l2.Next = l22
	l22.Val = 6
	l23 := new(ListNode)
	l22.Next = l23
	l23.Val = 4
	l23.Next = nil
	ans := addTwoNumbers(l1, l2)
	var got []int
	for ; ans != nil; ans = ans.Next {
		got = append(got, ans.Val)
	}
	want := []int{7, 0, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("addTwoNumbers(%v %v)=%v want:", l1, l2, got, want)
	}
}
