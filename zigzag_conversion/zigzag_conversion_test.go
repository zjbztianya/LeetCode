package zigzag_conversion

import "testing"

func TestConvert(t *testing.T) {
	var tests = []struct {
		s       string
		numRows int
		output  string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	}
	for i, test := range tests {
		if got := convert(test.s, test.numRows); got != test.output {
			t.Errorf("test case %d: output: %s ans: %s", i, test.output, got)
		}
	}
}
