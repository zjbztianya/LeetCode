package multiply_strings

import "testing"

func TestMultiply(t *testing.T) {
	var tests = []struct {
		num1   string
		num2   string
		output string
	}{{"0", "0", "0"},
		{"123", "12", "1476"},
		{"999", "999", "998001"},
		{"9999","0","0"},
	}

	for _, test := range tests {
		if got := multiply(test.num1, test.num2); got != test.output {
			t.Errorf("trap(%s %s)=%s want:%s", test.num1, test.num2, got, test.output)
		}
	}
}
