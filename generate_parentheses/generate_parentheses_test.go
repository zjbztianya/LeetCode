package generate_parentheses

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	var tests = []struct {
		input  int
		output []string
	}{
		{1, []string{"()"}},
		{2, []string{"()()", "(())"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for i, test := range tests {
		if got := generateParenthesis(test.input); !reflect.DeepEqual(got, test.output) {
			t.Errorf("test case %d: ans: %v want: %v", i, got, test.output)
		}
	}
}
