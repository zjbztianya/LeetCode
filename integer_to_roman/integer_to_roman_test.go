package integer_to_roman

import "testing"

func TestIntToRoman(t *testing.T) {
	var tests = []struct {
		input  int
		output string
	}{
		{1, "I"},
		{5, "V"},
		{199, "CXCIX"},
		{1437, "MCDXXXVII"},
		{1800, "MDCCC"},
		{3333, "MMMCCCXXXIII"},
	}

	for i, test := range tests {
		if got := intToRoman(test.input); got != test.output {
			t.Errorf("test case %d: output: %s ans: %s", i, test.output, got)
		}
	}
}
