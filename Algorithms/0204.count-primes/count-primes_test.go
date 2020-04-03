package problem0204

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n int
	ans int
}{



	// 可以有多个 testcase
}

func Test_countPrimes(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, countPrimes(tc.n), "输入:%v", tc)
	}
}

func Benchmark_countPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countPrimes(tc.n)
		}
	}
}
