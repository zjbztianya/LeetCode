package maximum_average_subarray

import (
	"testing"
	"math"
)

func TestFindMaxAverage(t *testing.T)  {
	var tests = struct {
		input []int
		k int
		want  float64
	}{
		[]int{1,12,-5,-6,50,3},4,12.75,
	}
	got:=findMaxAverage(tests.input,tests.k)
	if math.Abs(got-tests.want)>=0.01{
		t.Errorf("findMaxAverage(%v,%v)=%f", tests.input,tests.k, got)
	}
}