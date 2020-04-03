package problem0200

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	grid [][]byte
	ans int
}{



	// 可以有多个 testcase
}

func Test_numIslands(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, numIslands(tc.grid), "输入:%v", tc)
	}
}

func Benchmark_numIslands(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numIslands(tc.grid)
		}
	}
}
