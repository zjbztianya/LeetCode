package longest_common_prefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	var tests = []struct {
		input  []string
		output string
	}{
		{[]string{"aaa", "aa", "a"}, "a"},
		{[]string{"a", "b", "c"}, ""},
		{[]string{"ab", "abc", "abd", "ab"}, "ab"},
		{[]string{""}, ""},
	}

	for i, test := range tests {
		if got := longestCommonPrefix(test.input); got != test.output {
			t.Errorf("test case %d: output: %s ans: %s", i, test.output, got)
		}
	}
}
