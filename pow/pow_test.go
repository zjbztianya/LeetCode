package pow

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	var test = struct {
		x    float64
		n    int
		want float64
	}{8.88023, 3, 700.28148}

	if got := myPow(test.x, test.n); math.Abs(got-test.want) > 0.0001 {
		t.Errorf("myPow(%f %d)=%f want:%f", test.x, test.n, got, test.want)
	}
}
