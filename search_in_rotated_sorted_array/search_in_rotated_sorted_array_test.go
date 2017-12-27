package search_in_rotated_sorted_array

import "testing"

func TestSearch(t *testing.T) {
	var tests = []struct {
		input  []int
		target int
		want   int
	}{
		{[]int{}, 5, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 5},
		{[]int{4, 5, 6, 1, 2, 3}, 3, 6},
		{[]int{4, 5, 6, 1, 2, 3}, 4, 1},
	}

	for i, test := range tests {
		if got := search(test.input, test.target); got != test.want {
			t.Errorf("test case %d: output: %d want: %d", i, got, test.input)
		}
	}
}
