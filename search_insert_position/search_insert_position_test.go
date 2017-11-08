package search_insert_position

import "testing"

func TestSearchInsert(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		output int
	}{{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
	}

	for _, test := range tests {
		if got := searchInsert(test.nums, test.target); got != test.output {
			t.Errorf("searchInsert(%v %d)=%v want:%v", test.nums, test.target, got, test.output)
		}
	}
}
