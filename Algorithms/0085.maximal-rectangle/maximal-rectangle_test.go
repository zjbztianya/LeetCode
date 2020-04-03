package problem0085

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]byte
	ans    int
}{

	// 可以有多个 testcase
}

func Test_maximalRectangle(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, maximalRectangle(tc.matrix), "输入:%v", tc)
	}
}

func Benchmark_maximalRectangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maximalRectangle(tc.matrix)
		}
	}
}
