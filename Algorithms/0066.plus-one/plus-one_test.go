package problem0066

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	digits []int
	ans    []int
}{

	// 可以有多个 testcase
}

func Test_plusOne(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, plusOne(tc.digits), "输入:%v", tc)
	}
}

func Benchmark_plusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			plusOne(tc.digits)
		}
	}
}
