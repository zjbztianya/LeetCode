package find_substring

import (
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	var tests = []struct {
		s     string
		words []string
		want  []int
	}{
		{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
		{"aaaaa", []string{"b", "c"}, nil},
		{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
		{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
	}

	for i, test := range tests {
		if got := findSubstring(test.s, test.words); !reflect.DeepEqual(got, test.want) {
			t.Errorf("test case %d: output: %v want: %v", i, got, test.want)
		}
	}
}
