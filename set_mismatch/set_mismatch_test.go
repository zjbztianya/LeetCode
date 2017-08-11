package set_mismatch

import (
	"reflect"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 2, 4}, []int{2, 3}},
		{[]int{2, 2}, []int{2, 1}},
	}

	for _, test := range tests {
		if got := findErrorNums(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("countSubstrings(%q)=%v", test.input, got)
		}
	}
}
