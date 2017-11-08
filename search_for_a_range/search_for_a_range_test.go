package search_for_a_range

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		output []int
	}{{[]int{1}, 1, []int{0, 0}},
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{2,2},3,[]int{-1,-1}},
	}

	for _, test := range tests {
		if got := searchRange(test.nums, test.target); !reflect.DeepEqual(got, test.output) {
			t.Errorf("searchRange(%v %d)=%v want:%v", test.nums, test.target, got, test.output)
		}
	}
}
