package three_sum_closest

import "testing"

func TestThreeSumClosest(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		output int
	}{
		{[]int{0, 0, 0}, 1, 0},
		{[]int{-1, 2, 1, -4}, 1, 2},
	}

	for i, test := range tests {
		if got := threeSumClosest(test.nums, test.target); got != test.output {
			t.Errorf("test case %d: output: %d ans: %d", i, test.output, got)
		}
	}
}
