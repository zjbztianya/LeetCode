package longest_palindromic_substring

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T)  {
	var tests = []struct {
		input string
		output string
	}{
		{"babad","aba"},
		{"cbbd","bb"},
	}
	for i,test:=range tests{
		if got:=longestPalindrome(test.input); got!=test.output{
			t.Errorf("test case %d: output: %s ans: %s",i,test.output,got)
		}
	}
}
