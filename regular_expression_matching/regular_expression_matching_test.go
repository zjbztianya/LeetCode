package regular_expression_matching

import "testing"

func TestIsMatch(t *testing.T) {
	var tests = []struct {
		s      string
		p      string
		output bool
	}{
		{"aa", "a", false},
		{"aa", "aa", true},
		{"aaa", "aa", false},
		{"aa", "a*", true},
		{"aa", ".*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"ab", ".*c", false},
	}
	for i, test := range tests {
		if got := isMatch(test.s, test.p); got != test.output {
			t.Errorf("test case %d: output: %v ans: %v", i, test.output, got)
		}
	}
}
