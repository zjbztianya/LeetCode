package implement_strstr

import "testing"

func TestStrStr(t *testing.T) {
	var tests = []struct {
		haystack string
		needle   string
		output   int
	}{{"aaaa", "aa", 0},
		{"", "", 0},
		{"aabca", "ab", 1},
		{"aa", "", 0},
		{"", "aa", -1},
	}

	for _, test := range tests {
		if got := strStr(test.haystack, test.needle); got != test.output {
			t.Errorf("searchRange(%s %s)=%d want:%d", test.haystack, test.needle, got, test.output)
		}
	}
}
