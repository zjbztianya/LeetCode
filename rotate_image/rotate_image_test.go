package rotate_image

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	var test = struct {
		input  [][]int
		output [][]int
	}{
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
	}
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	if rotate(test.input); !reflect.DeepEqual(test.input, test.output) {
		t.Errorf("rotate(%v)=%v want:%v", input, test.input, test.output)
	}
}
