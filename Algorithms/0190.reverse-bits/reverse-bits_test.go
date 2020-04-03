package problem0190

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num uint32
	ans uint32
}{



	// 可以有多个 testcase
}

func Test_reverseBits(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, reverseBits(tc.num), "输入:%v", tc)
	}
}

func Benchmark_reverseBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverseBits(tc.num)
		}
	}
}
