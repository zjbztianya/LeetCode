package problem0023

import (
	"container/heap"
	"math"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *IntSlice) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *IntSlice) Pop() interface{} {
	top := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return top
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var a IntSlice
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		a = append(a, lists[i].Val)
		lists[i] = lists[i].Next
	}

	dummyHead := &ListNode{0, nil}
	p := dummyHead
	heap.Init(&a)

	for {
		var minV interface{}
		minV = math.MaxInt32
		idx := 0
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if minV.(int) > lists[i].Val {
				minV = lists[i].Val
				idx = i
			}
		}
		if minV == math.MaxInt32 {
			break
		}
		v := heap.Pop(&a)
		heap.Push(&a, minV)
		lists[idx] = lists[idx].Next
		q := &ListNode{v.(int), nil}
		p.Next = q
		p = p.Next
	}

	for len(a) > 0 {
		v := heap.Pop(&a)
		q := &ListNode{v.(int), nil}
		p.Next = q
		p = p.Next
	}

	return dummyHead.Next
}
