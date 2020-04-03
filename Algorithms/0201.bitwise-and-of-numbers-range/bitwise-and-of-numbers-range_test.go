package problem0201

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	m int
 n int
	ans int
}{



	// 可以有多个 testcase
}

func Test_rangeBitwiseAnd(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, rangeBitwiseAnd(tc.m, tc.n), "输入:%v", tc)
	}
}

func Benchmark_rangeBitwiseAnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			rangeBitwiseAnd(tc.m, tc.n)
		}
	}
}
