package next_permutation

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{1, 5, 1}, []int{5, 1, 1}},
		{[]int{5, 1, 1}, []int{1, 1, 5}},
	}

	for i, test := range tests {
		if nextPermutation(test.input); !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("test case %d: output: %v want: %v", i, test.input, test.want)
		}
	}
}
