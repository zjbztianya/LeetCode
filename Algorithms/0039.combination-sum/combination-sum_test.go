package problem0039

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	candidates []int
	target     int
	ans        [][]int
}{

	// 可以有多个 testcase
}

func Test_combinationSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, combinationSum(tc.candidates, tc.target), "输入:%v", tc)
	}
}

func Benchmark_combinationSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			combinationSum(tc.candidates, tc.target)
		}
	}
}
