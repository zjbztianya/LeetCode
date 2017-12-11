package three_sum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	var tests = []struct {
		input  []int
		output [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{[]int{2, 2, 2, -1, -1, -1}, [][]int{{2, -1, -1}}},
		{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}, [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}},
		{[]int{1, -1, -1, 0}, [][]int{{-1, 0, 1}}},
	}

	for i, test := range tests {
		if got := threeSum(test.input); !reflect.DeepEqual(got, test.output) {
			t.Errorf("test case %d: output: %v ans: %v", i, test.output, got)
		}
	}
}
