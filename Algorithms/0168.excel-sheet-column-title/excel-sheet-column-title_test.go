package problem0168

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans string
}{

	// 可以有多个 testcase
}

func Test_convertToTitle(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, convertToTitle(tc.n), "输入:%v", tc)
	}
}

func Benchmark_convertToTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			convertToTitle(tc.n)
		}
	}
}
