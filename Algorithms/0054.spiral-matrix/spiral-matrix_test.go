package problem0054

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]int
	ans    []int
}{

	// 可以有多个 testcase
}

func Test_spiralOrder(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, spiralOrder(tc.matrix), "输入:%v", tc)
	}
}

func Benchmark_spiralOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			spiralOrder(tc.matrix)
		}
	}
}
