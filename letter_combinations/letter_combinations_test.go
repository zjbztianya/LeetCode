package letter_combinations

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	var tests = []struct {
		input  string
		output []string
	}{
		{"", []string{}},
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}

	for i, test := range tests {
		if got := letterCombinations(test.input); !reflect.DeepEqual(test.output, got) {
			t.Errorf("test case %d: output: %v ans: %v", i, test.output, got)
		}
	}
}
