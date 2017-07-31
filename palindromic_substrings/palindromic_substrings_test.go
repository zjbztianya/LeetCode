package palindromic_substrings

import (
	"testing"
)

func TestCountSubstrings(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"abc", 3},
		{"aaa", 6},
	}

	for _, test := range tests {
		if got := countSubstrings(test.input); got != test.want {
			t.Errorf("countSubstrings(%q)=%v", test.input, got)
		}
	}

}
