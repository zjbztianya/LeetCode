package longest_valid_parentheses

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"(())", 4},
		{")))(", 0},
		{")()())", 4},
	}

	for i, test := range tests {
		if got := longestValidParentheses(test.input); got != test.want {
			t.Errorf("test case %d: output: %s want: %s", i, got, test.input)
		}
	}
}
