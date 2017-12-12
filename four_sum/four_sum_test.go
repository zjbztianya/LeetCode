package four_sum

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		output [][]int
	}{
		{[]int{0, 0, 0, 0}, 1, [][]int{}},
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}}},
	}

	for i, test := range tests {
		if got := fourSum(test.nums, test.target); !reflect.DeepEqual(got, test.output) {
			t.Errorf("test case %d: output: %v ans: %v", i, test.output, got)
		}
	}
}
