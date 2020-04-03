package problem0064

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	grid [][]int
	ans  int
}{

	// 可以有多个 testcase
}

func Test_minPathSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, minPathSum(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_minPathSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minPathSum(tc.grid)
		}
	}
}
