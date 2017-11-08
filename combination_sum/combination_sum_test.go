package combination_sum

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
		[]int{2, 3, 6, 7}, 7, [][]int{{7}, {2, 2, 3}},
	}
	if got := combinationSum(test.candidates, test.target); !reflect.DeepEqual(got, test.output) {
		t.Errorf("combinationSum(%v %d)=%v want:%v", test.candidates, test.target, got, test.output)
	}
}
