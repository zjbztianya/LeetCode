package longest_substring_without_repeating_characters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T)  {
	var tests = []struct{
		input  string
		output int
	}{
		{"abcabcbb", 3},
		{"bbbbb",1},
		{"pwwkew", 3},
		{"abc", 3},
		{"dvdf", 3},
		{"tmmzuxt", 5},
		{"cdd", 2},
	}

	for i:=range tests {
		if got:=lengthOfLongestSubstring(tests[i].input); got!=tests[i].output{
			t.Errorf("lengthOfLongestSubstring(%v)=%v",tests[i].input,got)
		}
	}
}
