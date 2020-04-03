package problem0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans string
}{

	// 可以有多个 testcase
}

func Test_longestPalindrome(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, longestPalindrome(tc.s), "输入:%v", tc)
	}
}

func Benchmark_longestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestPalindrome(tc.s)
		}
	}
}
