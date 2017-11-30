package median_of_two_sorted_arrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	var tests = []struct {
		num1   []int
		num2   []int
		output float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for i, test := range tests {
		if got := findMedianSortedArrays(test.num1, test.num2); got != test.output {
			t.Errorf("test case %d: output: %f ans: %f", i+1, got, test.output)
		}
	}
}
