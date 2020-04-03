package problem0130

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	board [][]byte
	ans
}{

	// 可以有多个 testcase
}

func Test_solve(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, solve(tc.board), "输入:%v", tc)
	}
}

func Benchmark_solve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			solve(tc.board)
		}
	}
}
