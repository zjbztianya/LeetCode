package reverse_integer

import "testing"

func TestConvert(t *testing.T) {
	var tests = []struct {
		input  int
		output int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
	}
	for i, test := range tests {
		if got := reverse(test.input); got != test.output {
			t.Errorf("test case %d: output: %d ans: %d", i, test.output, got)
		}
	}
}
