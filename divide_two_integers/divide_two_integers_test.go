package divide_two_integers

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	var tests = []struct {
		dividend int
		divisor  int
		want     int
	}{
		{0, 1, 0},
		{math.MinInt32, -1, math.MaxInt32},
		{100, -1, -100},
		{199, 3, 66},
		{-2147483648, -3, 715827882},
	}

	for i, test := range tests {
		if got := divide(test.dividend, test.divisor); got != test.want {
			t.Errorf("test case %d: output: %v want: %v", i, got, test.want)
		}
	}
}
