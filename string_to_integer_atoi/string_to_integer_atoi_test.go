package string_to_integer_atoi

import (
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	var tests = []struct {
		input  string
		output int
	}{
		{"  +123", 123},
		{"  -123", -123},
		{"1x11", 1},
		{"+2147483648", math.MaxInt32},
		{"-2147483649", math.MinInt32},
		{"0000", 0},
		{"-0000", 0},
		{"", 0},
	}
	for i, test := range tests {
		if got := myAtoi(test.input); got != test.output {
			t.Errorf("test case %d: output: %d ans: %d", i, test.output, got)
		}
	}
}
