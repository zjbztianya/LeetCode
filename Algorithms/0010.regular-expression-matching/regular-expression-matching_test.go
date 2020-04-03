package problem0010

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	p   string
	ans bool
}{

	// 可以有多个 testcase
}

func Test_isMatch(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, isMatch(tc.s, tc.p), "输入:%v", tc)
	}
}

func Benchmark_isMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isMatch(tc.s, tc.p)
		}
	}
}
