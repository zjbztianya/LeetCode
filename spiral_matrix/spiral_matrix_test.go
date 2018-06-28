package spiral_matrix

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	var tests = []struct {
		input [][]int
		want  []int
	}{{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}}}

	for i, test := range tests {
		if got := spiralOrder(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("test case %d: output: %d want: %d", i, got, test.want)
		}
	}
}
