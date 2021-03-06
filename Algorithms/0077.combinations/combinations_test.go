package problem0077

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	k   int
	ans [][]int
}{

	// 可以有多个 testcase
}

func Test_combine(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, combine(tc.n, tc.k), "输入:%v", tc)
	}
}

func Benchmark_combine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			combine(tc.n, tc.k)
		}
	}
}
