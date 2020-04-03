package problem0216

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	k int
 n int
	ans [][]int
}{



	// 可以有多个 testcase
}

func Test_combinationSum3(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, combinationSum3(tc.k, tc.n), "输入:%v", tc)
	}
}

func Benchmark_combinationSum3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			combinationSum3(tc.k, tc.n)
		}
	}
}
