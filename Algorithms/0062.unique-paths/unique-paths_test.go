package problem0062

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	m   int
	n   int
	ans int
}{

	// 可以有多个 testcase
}

func Test_uniquePaths(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, uniquePaths(tc.m, tc.n), "输入:%v", tc)
	}
}

func Benchmark_uniquePaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			uniquePaths(tc.m, tc.n)
		}
	}
}
