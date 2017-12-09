package roman_to_integer

import "testing"

func TestRomanToInt(t *testing.T) {
	var tests = []struct {
		output int
		input  string
	}{
		{1, "I"},
		{5, "V"},
		{199, "CXCIX"},
		{1437, "MCDXXXVII"},
		{1800, "MDCCC"},
		{3333, "MMMCCCXXXIII"},
	}

	for i, test := range tests {
		if got := romanToInt(test.input); got != test.output {
			t.Errorf("test case %d: output: %s ans: %s", i, test.output, got)
		}
	}

}
