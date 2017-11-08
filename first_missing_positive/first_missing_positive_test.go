package first_missing_positive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	var tests = []struct {
		input  []int
		output int
	}{{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{}, 1},
		{[]int{1},2},
	}

	for _, test := range tests {
		if got := firstMissingPositive(test.input); got != test.output {
			t.Errorf("firstMissingPositive(%v)=%d want:%d", test.input, got, test.output)
		}
	}
}
