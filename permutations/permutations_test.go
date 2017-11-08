package permutations

import (
	"testing"
	"reflect"
)

func TestPermute(t *testing.T) {
	var test = struct {
		input  []int
		output [][]int
	}{
		[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
	}
	if got:=permute(test.input); !reflect.DeepEqual(got,test.output){
		t.Errorf("permute(%v)=%v want:%v",test.input,got,test.output)
	}
}
