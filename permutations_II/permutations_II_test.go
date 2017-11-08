package permutations_II

import (
	"testing"
	"reflect"
)

func TestPermuteUnique(t *testing.T) {
	var test = struct {
		input  []int
		output [][]int
	}{
		[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
	}
	if got:=permuteUnique(test.input); !reflect.DeepEqual(got,test.output){
		t.Errorf("permuteUnique(%v)=%v want:%v",test.input,got,test.output)
	}
}
