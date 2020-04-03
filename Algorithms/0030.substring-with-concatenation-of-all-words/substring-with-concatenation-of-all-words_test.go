package problem0030

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s     string
	words []string
	ans   []int
}{

	// 可以有多个 testcase
}

func Test_findSubstring(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, findSubstring(tc.s, tc.words), "输入:%v", tc)
	}
}

func Benchmark_findSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findSubstring(tc.s, tc.words)
		}
	}
}
