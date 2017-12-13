package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	var tests = []struct {
		input  string
		output bool
	}{
		{"", true},
		{"()({})[]", true},
		{"()){}", false},
		{"{", false},
		{"([)]", false},
	}

	for i, test := range tests {
		if got := isValid(test.input); got != test.output {
			t.Errorf("test case %d: output: %v want: %v", i, test.output, got)
		}
	}
}
