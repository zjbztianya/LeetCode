package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for i := range tests {
		if got := twoSum(tests[i].nums, tests[i].target); !reflect.DeepEqual(got, tests[i].want) {
			t.Errorf("twoSum(%v %v)=%v, want:%v", tests[i].nums, tests[i].target, got, tests[i].want)
		}
	}
}
