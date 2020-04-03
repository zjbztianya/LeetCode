package problem0069

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	x   int
	ans int
}{

	// 可以有多个 testcase
}

func Test_mySqrt(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, mySqrt(tc.x), "输入:%v", tc)
	}
}

func Benchmark_mySqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			mySqrt(tc.x)
		}
	}
}
