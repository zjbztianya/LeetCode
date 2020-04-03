package problem0050

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	x   float64
	n   int
	ans float64
}{

	// 可以有多个 testcase
}

func Test_myPow(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, myPow(tc.x, tc.n), "输入:%v", tc)
	}
}

func Benchmark_myPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			myPow(tc.x, tc.n)
		}
	}
}
