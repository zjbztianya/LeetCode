package problem0029

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	dividend int
	divisor  int
	ans      int
}{

	// 可以有多个 testcase
}

func Test_divide(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, divide(tc.dividend, tc.divisor), "输入:%v", tc)
	}
}

func Benchmark_divide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			divide(tc.dividend, tc.divisor)
		}
	}
}
