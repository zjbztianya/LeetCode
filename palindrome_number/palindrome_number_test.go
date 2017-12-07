package palindrome_number

import (
	"math"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input  int
		output bool
	}{
		{-123, false},
		{121, true},
		{1221, true},
		{math.MaxInt32, false},
		{0, true},
	}
	for i, test := range tests {
		if got := isPalindrome(test.input); got != test.output {
			t.Errorf("test case %d: output: %d ans: %d", i, test.output, got)
		}
	}
}
