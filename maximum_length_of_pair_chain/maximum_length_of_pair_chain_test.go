package maximum_length_of_pair_chain

import "testing"

func TestFindLongestChain(t *testing.T) {
	var tests = []struct {
		input [][]int
		want  int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}}, 2},
	}
	for _, test := range tests {
		if got := findLongestChain(test.input); got != test.want {
			t.Errorf("findLongestChain(%q)=%v", test.input, got)
		}
	}
}
