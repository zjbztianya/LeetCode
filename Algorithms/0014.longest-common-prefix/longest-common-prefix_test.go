package problem0014

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	strs []string
	ans  string
}{

	// 可以有多个 testcase
}

func Test_longestCommonPrefix(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, longestCommonPrefix(tc.strs), "输入:%v", tc)
	}
}

func Benchmark_longestCommonPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestCommonPrefix(tc.strs)
		}
	}
}
