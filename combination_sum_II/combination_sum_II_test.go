package combination_sum_II

import (
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	var test = struct {
		candidates []int
		target     int
		output     [][]int
	}{
		[]int{10, 1, 2, 7, 6, 1, 5}, 8,
		[][]int{{1, 7}, {1, 2, 5}, {2, 6}, {1, 1, 6}},
	}
	if got := combinationSum2(test.candidates, test.target); !reflect.DeepEqual(got, test.output) {
		t.Errorf("combinationSum2(%v %d)=%v want:%v", test.candidates, test.target, got, test.output)
	}
}
