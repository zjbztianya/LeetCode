package problem0140

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s        string
	wordDict []string
	ans      []string
}{

	// 可以有多个 testcase
}

func Test_wordBreak(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, wordBreak(tc.s, tc.wordDict), "输入:%v", tc)
	}
}

func Benchmark_wordBreak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			wordBreak(tc.s, tc.wordDict)
		}
	}
}
