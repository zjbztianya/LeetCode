package group_anagrams

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	var test = struct {
		input  []string
		output [][]string
	}{[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
		[][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}},
	}

	if got := groupAnagrams(test.input); !reflect.DeepEqual(got, test.output) {
		t.Errorf("groupAnagrams(%v)=%v want:%v", test.input, got, test.output)
	}
}
