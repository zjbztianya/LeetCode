package problem0020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans bool
}{

	// 可以有多个 testcase
}

func Test_isValid(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, isValid(tc.s), "输入:%v", tc)
	}
}

func Benchmark_isValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isValid(tc.s)
		}
	}
}
