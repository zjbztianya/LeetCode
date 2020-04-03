package problem0011

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	height []int
	ans    int
}{

	// 可以有多个 testcase
}

func Test_maxArea(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, maxArea(tc.height), "输入:%v", tc)
	}
}

func Benchmark_maxArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxArea(tc.height)
		}
	}
}
