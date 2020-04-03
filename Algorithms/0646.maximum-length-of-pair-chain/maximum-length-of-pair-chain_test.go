package problem0646

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pairs [][]int
	ans int
}{



	// 可以有多个 testcase
}

func Test_findLongestChain(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, findLongestChain(tc.pairs), "输入:%v", tc)
	}
}

func Benchmark_findLongestChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findLongestChain(tc.pairs)
		}
	}
}
