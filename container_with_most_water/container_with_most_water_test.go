package container_with_most_water

import "testing"

func TestMaxArea(t *testing.T) {
	var tests = []struct {
		input  []int
		output int
	}{
		{[]int{3, 1, 3}, 6},
		{[]int{1, 1}, 1},
		{[]int{1, 20, 20, 1}, 20},
	}

	for i, test := range tests {
		if got := maxArea(test.input); got != test.output {
			t.Errorf("test case %d: output: %d ans: %d", i, test.output, got)
		}
	}
}
